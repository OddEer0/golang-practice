package pgxUseCase

import (
	"context"
	"github.com/OddEer0/golang-practice/resources/sql"
	"log/slog"
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
		Value    PureUser
		Posts    []*model.Post
		Comments []*model.Comment
	}

	UserUseCase interface {
		Create(context.Context, *CreateUserData) (PureUser, error)
		GetUserById(context.Context, domain.Id, model.UserConns) (PureUserAggregate, error)
		GetUserByQuery(context.Context, domain.Id, *model.ManyOpt, model.UserConns) ([]*PureUserAggregate, error)
		UpdateUserLogin(context.Context, domain.Id, string) (PureUser, error)
	}

	userUseCase struct {
		userRepository    repository.User
		postRepository    repository.Post
		commentRepository repository.Comment
		transactor        sql.Transactor
	}
)

func (u *userUseCase) GetUserByQuery(ctx context.Context, id domain.Id, opt *model.ManyOpt, conns model.UserConns) ([]*PureUserAggregate, error) {
	users, err := u.userRepository.GetByQuery(ctx, opt)
	if err != nil {
		return nil, err
	}
	userAggregates := make([]*PureUserAggregate, 0, len(users))
	for _, user := range users {
		pureUserAggregate := &PureUserAggregate{
			Value: PureUser{
				Id:    user.Id,
				Login: user.Login,
				Email: user.Email,
			},
		}

		if option, ok := conns["posts"]; ok {
			posts, err := u.postRepository.GetByOwnerId(ctx, user.Id, option)
			if err != nil {
				return nil, err
			}
			pureUserAggregate.Posts = posts
		}

		if option, ok := conns["comments"]; ok {
			comments, err := u.commentRepository.GetByOwnerId(ctx, user.Id, option)
			if err != nil {
				return nil, err
			}
			pureUserAggregate.Comments = comments
		}

		userAggregates = append(userAggregates, pureUserAggregate)
	}

	return userAggregates, nil
}

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
	err = u.transactor.WithinTransaction(ctx, func(transactionCtx context.Context) error {
		user, err = u.userRepository.UpdateLoginById(transactionCtx, id, newLogin)
		if err != nil {
			return err
		}
		slog.Info("update user login success")
		_, err = u.postRepository.UpdateBodyByUserId(transactionCtx, &model.User{
			Id:    id,
			Login: newLogin,
		})
		if err != nil {
			return err
		}
		slog.Info("update posts body success")
		err = u.commentRepository.UpdateBodyByUserId(transactionCtx, &model.User{
			Id:    id,
			Login: newLogin,
		})
		if err != nil {
			return err
		}
		slog.Info("update comments body success")
		//return errors.New("check transaction work")
		return nil
	})

	if err != nil {
		return PureUser{}, err
	}

	return PureUser{
		Id:    user.Id,
		Login: user.Login,
		Email: user.Email,
	}, nil
}

func NewUserUseCase(userRepository repository.User, postRepository repository.Post, commentRepository repository.Comment, transactor sql.Transactor) UserUseCase {
	return &userUseCase{
		userRepository:    userRepository,
		postRepository:    postRepository,
		commentRepository: commentRepository,
		transactor:        transactor,
	}
}
