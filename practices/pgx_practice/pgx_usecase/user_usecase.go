package pgxUseCase

import (
	"context"
	"github.com/OddEer0/golang-practice/resources/sql"
	"time"

	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/google/uuid"
)

type (
	CreateUserData struct {
		Login    string
		Password string
		Email    string
	}

	PureUser struct {
		Id    domain.Id
		Login string
		Email string
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
		transactor     sql.Transactor
	}
)

func (u *userUseCase) Create(ctx context.Context, data *CreateUserData) (PureUser, error) {
	id := domain.Id(uuid.New().String())
	_, err := u.userRepository.Create(ctx, &model.User{
		Id:        id,
		Login:     data.Login,
		Email:     data.Email,
		Password:  data.Password, // FIXME - use bcrypt to hash password
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return PureUser{}, err
	}
	return PureUser{
		Id:    id,
		Login: data.Login,
		Email: data.Email,
	}, nil
}

// TODO - goroutines multi query
func (u *userUseCase) GetUserById(ctx context.Context, id domain.Id, conns model.UserConns) (PureUserAggregate, error) {
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
			Id:    user.Id,
			Login: user.Login,
			Email: user.Email,
		},
		Posts: posts,
	}, nil
}

func (u *userUseCase) UpdateUserLogin(ctx context.Context, id domain.Id, newLogin string) (PureUser, error) {
	user := &model.User{}
	var err error
	err := u.transactor.WithinTransaction(ctx, func(transactionCtx context.Context) error {
		user, err = u.userRepository.UpdateLoginById(ctx, id, newLogin)
		if err != nil {
			return err
		}
		_, err = u.postRepository.UpdateBodyByUserId(ctx, &model.User{
			Id:    id,
			Login: newLogin,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return PureUser{
		Id:    user.Id,
		Login: user.Login,
		Email: user.Email,
	}, nil
}

func NewUserUseCase(userRepository repository.User, postRepository repository.Post, transactor sql.Transactor) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
		postRepository: postRepository,
		transactor:     transactor,
	}
}
