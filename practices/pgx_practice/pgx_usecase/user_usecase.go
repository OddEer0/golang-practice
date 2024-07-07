package pgxUseCase

import (
	"context"

	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
)

type (
	CreateUserData struct {
		Login string
		Password string
		Email string
	}

	PureUser struct {
	}

	PureUserAggregate struct {
		Value PureUser
		Posts []*model.Post
	}

	UserUseCase interface {
		Create(context.Context, *CreateUserData) (PureUser, error)
		GetUserById(context.Context, domain.Id, model.UserConns) (PureUserAggregate, error)
		UpdateUserLogin(context.Context, domain.Id, string) (PureUser, error)
	}

	userUseCase struct {
		userRepository repository.User
		postRepository repository.Post
	}
)

func (u *userUseCase) Create(context.Context, *CreateUserData) (PureUser, error) {
	panic("unimplemented")
}

func (u *userUseCase) GetUserById(context.Context, domain.Id, model.UserConns) (PureUserAggregate, error) {
	panic("unimplemented")
}

func (u *userUseCase) UpdateUserLogin(context.Context, domain.Id, string) (PureUser, error) {
	panic("unimplemented")
}

func NewUserUseCase(userRepository repository.User, postRepository repository.Post) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
		postRepository: postRepository,
	}
}
