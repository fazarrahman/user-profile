// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"

	errorlib "github.com/fazarrahman/user-profile/errorLib"
)

type RepositoryInterface interface {
	CreateUser(ctx context.Context, input CreateUserInput) (output *CreateUserOutput, err *errorlib.Error)
	GetUserByPhoneNumber(ctx context.Context, input GetUserByPhoneNumberInput) (*Users, *errorlib.Error)
	UpdateSuccessfulLoginCount(ctx context.Context, input UpdateSuccessfulLoginCountInput) *errorlib.Error
	GetUserById(ctx context.Context, input GetUserByIdInput) (*Users, *errorlib.Error)
}
