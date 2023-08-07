package service

import (
	"context"
	"net/http"
	"testing"

	"github.com/fazarrahman/user-profile/generated"
	"github.com/fazarrahman/user-profile/repository"
	"github.com/golang/mock/gomock"
)

func TestCreateUser_PhoneNbrNotProvided_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var fullName string = "aaaa"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.CreateUser(c, generated.Users{
		FullName: &fullName,
	})
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Phone numbers is required" {
		t.Fail()
	}
}

func TestCreateUser_PhoneNbrIncorrectLength_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var fullName string = "aaaa"
	var phoneNbr string = "+628"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.CreateUser(c, generated.Users{
		FullName:    &fullName,
		PhoneNumber: &phoneNbr,
	})
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Phone numbers must be at minimum 10 characters and maximum 13 characters" {
		t.Fail()
	}
}

func TestCreateUser_PhoneNbrDoesntContainIndonesianPhoneCode_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var fullName string = "aaaa"
	var phoneNbr string = "08124378944"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.CreateUser(c, generated.Users{
		FullName:    &fullName,
		PhoneNumber: &phoneNbr,
	})
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Phone numbers must start with the Indonesia country code +62" {
		t.Fail()
	}
}

func TestCreateUser_PhoneNbrAlreadyUsed_ReturnConflictStatusCode(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var fullName string = "aaaa"
	var phoneNbr string = "+628123352342"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)
	repoMock.EXPECT().GetUserByPhoneNumber(c,
		repository.GetUserByPhoneNumberInput{PhoneNumber: phoneNbr}).Return(&repository.Users{Id: 1,
		PhoneNumber: phoneNbr}, nil)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.CreateUser(c, generated.Users{
		FullName:    &fullName,
		PhoneNumber: &phoneNbr,
	})
	if err.StatusCode != http.StatusConflict ||
		err.Message != "Phone number is already used" {
		t.Fail()
	}
}

func TestCreateUser_FullNameIsNotProvided_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var phoneNbr string = "+628123352342"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)
	repoMock.EXPECT().GetUserByPhoneNumber(c,
		repository.GetUserByPhoneNumberInput{PhoneNumber: phoneNbr}).Return(nil, nil)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.CreateUser(c, generated.Users{
		PhoneNumber: &phoneNbr,
	})
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Full name is required" {
		t.Fail()
	}
}

func TestCreateUser_FullNameIncorrectLength_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var fullName string = "a"
	var phoneNbr string = "+628123352342"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)
	repoMock.EXPECT().GetUserByPhoneNumber(c,
		repository.GetUserByPhoneNumberInput{PhoneNumber: phoneNbr}).Return(nil, nil)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.CreateUser(c, generated.Users{
		FullName:    &fullName,
		PhoneNumber: &phoneNbr,
	})
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Full name must be at minimum 3 characters and maximum 60 characters" {
		t.Fail()
	}
}

func TestCreateUser_PasswordNotProvided_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var fullName string = "Fazar Rahman"
	var phoneNbr string = "+628123352342"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)
	repoMock.EXPECT().GetUserByPhoneNumber(c,
		repository.GetUserByPhoneNumberInput{PhoneNumber: phoneNbr}).Return(nil, nil)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.CreateUser(c, generated.Users{
		FullName:    &fullName,
		PhoneNumber: &phoneNbr,
	})
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Password is required" {
		t.Fail()
	}
}

func TestCreateUser_PasswordIncorrectLength_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var fullName string = "Fazar Rahman"
	var phoneNbr string = "+628123352342"
	var password string = "a"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)
	repoMock.EXPECT().GetUserByPhoneNumber(c,
		repository.GetUserByPhoneNumberInput{PhoneNumber: phoneNbr}).Return(nil, nil)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.CreateUser(c, generated.Users{
		FullName:    &fullName,
		PhoneNumber: &phoneNbr,
		Passwords:   &password,
	})
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Passwords must be minimum 6 characters and maximum 64 characters" {
		t.Fail()
	}
}

func TestCreateUser_CorrectPayload_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var fullName string = "Fazar Rahman"
	var phoneNbr string = "+628123352342"
	var password string = "Passwords"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)
	repoMock.EXPECT().GetUserByPhoneNumber(c,
		repository.GetUserByPhoneNumberInput{PhoneNumber: phoneNbr}).Return(nil, nil)
	repoMock.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&repository.CreateUserOutput{Id: 1}, nil)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.CreateUser(c, generated.Users{
		FullName:    &fullName,
		PhoneNumber: &phoneNbr,
		Passwords:   &password,
	})
	if err != nil {
		t.Fail()
	}
}

func TestLogin_PhoneNbrNotProvided_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var password string = "Passwords"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.Login(c, &generated.LoginInput{
		Passwords: &password,
	})
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Phone numbers is required" {
		t.Fail()
	}
}

func TestLogin_PasswordNotProvided_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var phoneNbr string = "+62834378293"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.Login(c, &generated.LoginInput{
		PhoneNumber: &phoneNbr,
	})
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Password is required" {
		t.Fail()
	}
}

func TestLogin_InvalidPhoneNbr_ReturnNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var phoneNbr string = "+62834378293"
	var password string = "Passwords"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)
	repoMock.EXPECT().GetUserByPhoneNumber(c,
		repository.GetUserByPhoneNumberInput{PhoneNumber: phoneNbr}).Return(nil, nil)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.Login(c, &generated.LoginInput{
		PhoneNumber: &phoneNbr,
		Passwords:   &password,
	})
	if err.StatusCode != http.StatusNotFound ||
		err.Message != "Invalid phone number" {
		t.Fail()
	}
}

func TestLogin_InvalidPassword_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()
	var phoneNbr string = "+62834378293"
	var password string = "Passwords"

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)
	repoMock.EXPECT().GetUserByPhoneNumber(c,
		repository.GetUserByPhoneNumberInput{PhoneNumber: phoneNbr}).Return(&repository.Users{Id: 1,
		PhoneNumber: phoneNbr, Password: []byte("Passwords2")}, nil)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.Login(c, &generated.LoginInput{
		PhoneNumber: &phoneNbr,
		Passwords:   &password,
	})
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Invalid password" {
		t.Fail()
	}
}

func TestGetUserByAccessToken_AccessTokenNotProvided_ReturnForbidden(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.GetUserByAccessToken(c, "")
	if err.StatusCode != http.StatusForbidden ||
		err.Message != "Access token is required" {
		t.Fail()
	}
}

func TestGetUserByAccessToken_InvalidAccessToken_ReturnForbidden(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	_, err := svc.GetUserByAccessToken(c, "aaa")
	if err.StatusCode != http.StatusForbidden {
		t.Fail()
	}
}

func TestUpdateUser_PhoneNbrNotProvided_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	err := svc.UpdateUser(c, generated.UpdateUsers{}, "aaa")
	if err.StatusCode != http.StatusBadRequest || err.Message != "Phone numbers is required" {
		t.Fail()
	}
}

func TestUpdateUser_PhoneNbrIncorrectLength_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	var phoneNbr string = "+62812"
	err := svc.UpdateUser(c, generated.UpdateUsers{PhoneNumber: &phoneNbr}, "aaa")
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Phone numbers must be at minimum 10 characters and maximum 13 characters" {
		t.Fail()
	}
}

func TestUpdateUser_PhoneNbrDoesntStartWithIndonesiaCountryCode_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	var phoneNbr string = "081232449324"
	err := svc.UpdateUser(c, generated.UpdateUsers{PhoneNumber: &phoneNbr}, "aaa")
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Phone numbers must start with the Indonesia country code +62" {
		t.Fail()
	}
}

func TestUpdateUser_FullNameNotProvided_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	var phoneNbr string = "+6281232449"
	err := svc.UpdateUser(c, generated.UpdateUsers{PhoneNumber: &phoneNbr}, "aaa")
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Full name is required" {
		t.Fail()
	}
}

func TestUpdateUser_FullNameIncorrectLength_ReturnBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c := context.Background()

	repoMock := repository.NewMockRepositoryInterface(mockCtrl)

	svc := NewService(NewServiceOptions{
		Repository: repoMock,
	})

	var phoneNbr string = "+6281232449"
	var fullName string = "an"
	err := svc.UpdateUser(c, generated.UpdateUsers{PhoneNumber: &phoneNbr, FullName: &fullName}, "aaa")
	if err.StatusCode != http.StatusBadRequest ||
		err.Message != "Full name must be at minimum 3 characters and maximum 60 characters" {
		t.Fail()
	}
}
