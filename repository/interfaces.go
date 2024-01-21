// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
)

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	IsPhoneNumberExists(ctx context.Context, phoneNumber string) (isExists bool, err error)
	CreateUser(ctx context.Context, input CreateUserInput) (output *CreateUserOutput, err error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (output *UserOutput, err error)
	GetUserById(ctx context.Context, id uint) (output *UserOutput, err error)
	IncrementUserLoginCount(ctx context.Context, id uint) error
}
