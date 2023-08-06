package service

import (
	"context"
	"strings"

	errorlib "github.com/fazarrahman/user-profile/errorLib"
	"github.com/fazarrahman/user-profile/generated"
	"github.com/fazarrahman/user-profile/repository"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateUser(ctx context.Context, req generated.Users) (output *repository.CreateUserOutput, errl *errorlib.Error) {
	if req.PhoneNumbers == nil {
		return nil, errorlib.BadRequest("Phone numbers is required")
	}
	if len(*req.PhoneNumbers) < 10 || len(*req.PhoneNumbers) > 13 {
		return nil, errorlib.BadRequest("Phone numbers must be at minimum 10 characters and maximum 13 characters")
	}
	pArr := strings.Split(*req.PhoneNumbers, "")
	if strings.Join(pArr[:3], "") != "+62" {
		return nil, errorlib.BadRequest("Phone numbers must start with the Indonesia country code +62")
	}
	userByPhone, errl := s.Repository.GetUserByPhoneNumber(ctx,
		repository.GetUserByPhoneNumberInput{PhoneNumber: *req.PhoneNumbers})
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
		PhoneNumber: *req.PhoneNumbers,
		FullName:    *req.FullName,
		Password:    pwdhash,
	})
}
