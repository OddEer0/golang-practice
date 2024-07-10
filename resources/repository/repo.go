package repository

import (
	"context"

	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
)

type (
	User interface {
		GetById(context.Context, domain.Id) (*model.User, error)
		Create(context.Context, *model.User) (*model.User, error)
		UpdateById(context.Context, *model.User) (*model.User, error)
		UpdateLoginById(context.Context, domain.Id, string) (*model.User, error)
		DeleteById(context.Context, domain.Id) error
	}

	Post interface {
		GetById(context.Context, domain.Id) (*model.Post, error)
		GetByOwnerId(context.Context, domain.Id, *model.ManyOpt) ([]*model.Post, error)
		Create(context.Context, *model.Post) (*model.Post, error)
		UpdateById(context.Context, *model.Post) (*model.Post, error)
		DeleteById(context.Context, domain.Id) error
		UpdateBodyByUserId(context.Context, *model.User) (*model.Post, error)
	}

	Comment interface {
		GetById(context.Context, domain.Id) (*model.Comment, error)
		GetByPostId(context.Context, domain.Id, *model.ManyOpt) ([]*model.Comment, error)
		Create(context.Context, *model.Comment) (*model.Comment, error)
		UpdateById(context.Context, *model.Comment) (*model.Comment, error)
		DeleteById(context.Context, domain.Id) error
		UpdateBodyByUserId(context.Context, *model.User) error
	}
)
