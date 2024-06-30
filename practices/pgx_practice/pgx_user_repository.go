package pgxPractice

import (
	"context"

	resourceModel "github.com/OddEer0/golang-practice/resources/resource_model"
	resourceRepository "github.com/OddEer0/golang-practice/resources/resource_repository"
)

type userRepository struct {
	db QueryExecutor
}

// Create implements resourceRepository.UserRepository.
func (u *userRepository) Create(context.Context, resourceModel.User) (*resourceModel.User, error) {
	panic("unimplemented")
}

// CreateCopy implements resourceRepository.UserRepository.
func (u *userRepository) CreateCopy(context.Context, resourceModel.User) (resourceModel.User, error) {
	panic("unimplemented")
}

// DeleteById implements resourceRepository.UserRepository.
func (u *userRepository) DeleteById(context.Context, string) error {
	panic("unimplemented")
}

// GetById implements resourceRepository.UserRepository.
func (u *userRepository) GetById(context.Context, string) (*resourceModel.User, error) {
	panic("unimplemented")
}

// GetByIdCopy implements resourceRepository.UserRepository.
func (u *userRepository) GetByIdCopy(context.Context, string) (resourceModel.User, error) {
	panic("unimplemented")
}

func NewUserRepository(conn QueryExecutor) resourceRepository.User {
	return &userRepository{
		db: conn,
	}
}
