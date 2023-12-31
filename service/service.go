package service

import (
	"context"

	errorlib "github.com/fazarrahman/user-profile/errorLib"
	"github.com/fazarrahman/user-profile/generated"
	"github.com/fazarrahman/user-profile/repository"
)

type Service struct {
	Repository repository.RepositoryInterface
}

type NewServiceOptions struct {
	Repository repository.RepositoryInterface
}

func NewService(opts NewServiceOptions) *Service {
	return &Service{Repository: opts.Repository}
}

type ServiceInterface interface {
	CreateUser(ctx context.Context, input generated.Users) (output *repository.CreateUserOutput, errl *errorlib.Error)
	Login(ctx context.Context, loginInput *generated.LoginInput) (*generated.LoginResponse, *errorlib.Error)
	GetUserByAccessToken(ctx context.Context, accessToken string) (*repository.Users, *errorlib.Error)
	UpdateUser(ctx context.Context, req generated.UpdateUsers, accessToken string) *errorlib.Error
}
