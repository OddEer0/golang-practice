package pgxpractice

import (
	"context"

	resourcemodel "github.com/OddEer0/golang-practice/resources/resource_model"
	resourcerepository "github.com/OddEer0/golang-practice/resources/resource_repository"
)

type userRepository struct {
	db Query
}

// Create implements resourcerepository.UserRepository.
func (u *userRepository) Create(context.Context, resourcemodel.User) (*resourcemodel.User, error) {
	panic("unimplemented")
}

// CreateCopy implements resourcerepository.UserRepository.
func (u *userRepository) CreateCopy(context.Context, resourcemodel.User) (resourcemodel.User, error) {
	panic("unimplemented")
}

// DeleteById implements resourcerepository.UserRepository.
func (u *userRepository) DeleteById(context.Context, string) error {
	panic("unimplemented")
}

// GetById implements resourcerepository.UserRepository.
func (u *userRepository) GetById(context.Context, string) (*resourcemodel.User, error) {
	panic("unimplemented")
}

// GetByIdCopy implements resourcerepository.UserRepository.
func (u *userRepository) GetByIdCopy(context.Context, string) (resourcemodel.User, error) {
	panic("unimplemented")
}

func NewUserRepository(conn Query) resourcerepository.UserRepository {
	return &userRepository{
		db: conn,
	}
}
