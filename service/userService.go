package service

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	errorlib "github.com/fazarrahman/user-profile/errorLib"
	"github.com/fazarrahman/user-profile/generated"
	"github.com/fazarrahman/user-profile/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateUser(ctx context.Context, req generated.Users) (output *repository.CreateUserOutput, errl *errorlib.Error) {
	if req.PhoneNumber == nil {
		return nil, errorlib.BadRequest("Phone numbers is required")
	}
	if len(*req.PhoneNumber) < 10 || len(*req.PhoneNumber) > 13 {
		return nil, errorlib.BadRequest("Phone numbers must be at minimum 10 characters and maximum 13 characters")
	}
	pArr := strings.Split(*req.PhoneNumber, "")
	if strings.Join(pArr[:3], "") != "+62" {
		return nil, errorlib.BadRequest("Phone numbers must start with the Indonesia country code +62")
	}
	userByPhone, errl := s.Repository.GetUserByPhoneNumber(ctx,
		repository.GetUserByPhoneNumberInput{PhoneNumber: *req.PhoneNumber})
	if errl != nil {
		return nil, errl
	}
	if userByPhone != nil {
		return nil, errorlib.ResourceAlreadyExist("Phone number is already exists")
	}
	if req.FullName == nil {
		return nil, errorlib.BadRequest("Full name is required")
	}
	if len(*req.FullName) < 3 || len(*req.FullName) > 60 {
		return nil, errorlib.BadRequest("Full name must be at minimum 3 characters and maximum 60 characters")
	}
	if req.Passwords == nil {
		return nil, errorlib.BadRequest("Password is required")
	}
	if len(*req.Passwords) < 6 || len(*req.Passwords) > 64 {
		return nil, errorlib.BadRequest("Passwords must be minimum 6 characters and maximum 64 characters")
	}
	pwdhash, err := bcrypt.GenerateFromPassword([]byte(*req.Passwords), bcrypt.DefaultCost)
	if err != nil {
		return nil, errorlib.InternalServerError("Error when encrpting password " + err.Error())
	}

	return s.Repository.CreateUser(ctx, repository.CreateUserInput{
		PhoneNumber: *req.PhoneNumber,
		FullName:    *req.FullName,
		Password:    pwdhash,
	})
}

func (s *Service) Login(ctx context.Context, loginInput *generated.LoginInput) (*generated.LoginResponse, *errorlib.Error) {
	if loginInput.PhoneNumber == nil {
		return nil, errorlib.BadRequest("Phone numbers is required")
	}
	if loginInput.Passwords == nil {
		return nil, errorlib.BadRequest("Password is required")
	}

	user, err := s.CheckPhoneNbrAndPassword(ctx, loginInput)
	if user == nil && err != nil {
		return nil, err
	}

	token, erro := GenerateToken(user)

	if erro != nil {
		return nil, errorlib.InternalServerError("Error when generating token : " + erro.Error())
	}

	s.Repository.UpdateSuccessfulLoginCount(ctx, repository.UpdateSuccessfulLoginCountInput{Id: user.Id})

	return &generated.LoginResponse{Id: user.Id, AccessToken: token}, nil

}

// CheckUsernamePassword ..
func (s *Service) CheckPhoneNbrAndPassword(ctx context.Context, r *generated.LoginInput) (*repository.Users, *errorlib.Error) {
	userEntity, err := s.Repository.GetUserByPhoneNumber(ctx, repository.GetUserByPhoneNumberInput{
		PhoneNumber: *r.PhoneNumber,
	})
	if userEntity == nil && err == nil {
		return nil, errorlib.NotFound("Invalid phone number")
	} else if err != nil {
		return nil, err
	}

	erro := bcrypt.CompareHashAndPassword(userEntity.Password, []byte(*r.Passwords))
	if erro != nil {
		return nil, errorlib.BadRequest("Invalid password")
	}
	return userEntity, nil
}

func (s *Service) GetUserByAccessToken(ctx context.Context, accessToken string) (*repository.Users, *errorlib.Error) {
	if accessToken == "" {
		return nil, errorlib.Forbidden("Access token is required")
	}
	claimMap, err := ExtractTokenID(ctx, accessToken)
	if err != nil {
		return nil, errorlib.Forbidden("Invalid access token : " + err.Error())
	}
	val := claimMap["Id"]
	idfloat := val.(float64)
	id := int64(idfloat)
	user, errl := s.Repository.GetUserById(ctx, repository.GetUserByIdInput{
		Id: id,
	})
	if errl != nil {
		return nil, errl
	}
	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, req generated.UpdateUsers, accessToken string) *errorlib.Error {
	user, errl := s.GetUserByAccessToken(ctx, accessToken)
	if errl != nil {
		return errl
	}
	var phoneNumber string = user.PhoneNumber
	var fullName string = user.FullName
	if req.PhoneNumber != nil && user.PhoneNumber != *req.PhoneNumber {
		userByNewPhone, errl := s.Repository.GetUserByPhoneNumber(ctx, repository.GetUserByPhoneNumberInput{
			PhoneNumber: *req.PhoneNumber,
		})
		if errl != nil {
			return errl
		}
		if userByNewPhone != nil {
			return errorlib.ResourceAlreadyExist("The phone number is already used")
		}
		phoneNumber = *req.PhoneNumber
	}
	if req.FullName != nil {
		fullName = *req.FullName
	}
	errl = s.Repository.UpdateUser(ctx, repository.UpdateUserInput{
		Id:          user.Id,
		PhoneNumber: phoneNumber,
		FullName:    fullName,
	})
	if errl != nil {
		return errl
	}
	return nil
}

func GenerateToken(user *repository.Users) (string, error) {

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["Id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func ExtractTokenID(c context.Context, tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var cMap = make(map[string]interface{})
		val := claims["Id"]
		cMap["Id"] = val.(float64)
		cMap["exp"] = fmt.Sprintf("%0.f", claims["exp"])
		return cMap, nil
	}
	return nil, nil
}
