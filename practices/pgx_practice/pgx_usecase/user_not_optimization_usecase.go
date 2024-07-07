package pgxUseCase

import (
	"context"
	"time"

	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/google/uuid"
)

type (
	UserNOUseCase interface {
		Create(context.Context, *CreateUserData) (PureUser, error)
		GetUserById(context.Context, domain.Id, model.UserConns) (PureUserAggregate, error)
		UpdateUserLogin(context.Context, domain.Id, string) (PureUser, error)
	}

	userNOUseCase struct {
		userRepository repository.User
		postRepository repository.Post
	}
)

func (u *userNOUseCase) Create(ctx context.Context, data *CreateUserData) (PureUser, error) {
	id := domain.Id(uuid.New().String())
	_, err := u.userRepository.Create(ctx, &model.User{
		Id: id,
		Login: data.Login,
		Email: data.Email,
		Password: data.Password, // FIXME - use bcrypt to hash password
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return PureUser{}, err
	}
	return PureUser{
		Id: id,
		Login: data.Login,
		Email: data.Email,
	}, nil
}


func (u *userNOUseCase) GetUserById(ctx context.Context, id domain.Id, conns model.UserConns) (PureUserAggregate, error) {
	user, err := u.userRepository.GetById(ctx, id)
	if err != nil {
		return PureUserAggregate{}, nil
	}

	var posts []*model.Post
	if opt, ok := conns["posts"]; ok {
		posts, err = u.postRepository.GetByOwnerId(ctx, id, opt)
		if err != nil {
			return PureUserAggregate{}, nil
		}
	}

	return PureUserAggregate{
		Value: PureUser{
			Id: user.Id,
			Login: user.Login,
			Email: user.Email,
		},
		Posts: posts,
	}, nil
}

// TODO - fix transaction
func (u *userNOUseCase) UpdateUserLogin(ctx context.Context, id domain.Id, newLogin string) (PureUser, error) {
	user, err := u.userRepository.UpdateLoginById(ctx, id, newLogin)
	if err != nil {
		return PureUser{}, err
	}

	_, err = u.postRepository.UpdateBodyByUserId(ctx, &model.User{
		Id: id,
		Login: newLogin,
	})
	if err != nil {
		return PureUser{}, err
	}

	return PureUser{
		Id: user.Id,
		Login: user.Login,
		Email: user.Email,
	}, nil
}

func NewUserNOUseCase(userRepository repository.User, postRepository repository.Post) UserUseCase {
	return &userNOUseCase{
		userRepository: userRepository,
		postRepository: postRepository,
	}
}
