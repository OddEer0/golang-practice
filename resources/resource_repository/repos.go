package resourcerepository

import (
	"context"

	resourcemodel "github.com/OddEer0/golang-practice/resources/resource_model"
)

type (
	UserRepository interface {
		GetById(context.Context, string) (*resourcemodel.User, error)
		GetByIdCopy(context.Context, string) (resourcemodel.User, error)
		Create(context.Context, resourcemodel.User) (*resourcemodel.User, error)
		CreateCopy(context.Context, resourcemodel.User) (resourcemodel.User, error)
		DeleteById(context.Context, string) error
	}
	PostRepository interface {
		GetById(context.Context, string) (*resourcemodel.Post, error)
		GetByIdCopy(context.Context, string) (resourcemodel.Post, error)
		Create(context.Context, resourcemodel.Post) (*resourcemodel.Post, error)
		CreateCopy(context.Context, resourcemodel.Post) (resourcemodel.Post, error)
		DeleteById(context.Context, string) error

		GetPostAggregate(context.Context, string, *resourcemodel.AggregateOption) (*resourcemodel.PostAggregate, error)
		GetPostAggregateCopy(context.Context, string, *resourcemodel.AggregateOption) (resourcemodel.PostAggregate, error)
	}
	CommentRepository interface {
		GetById(context.Context, string) (*resourcemodel.Comment, error)
		GetByIdCopy(context.Context, string) (resourcemodel.Comment, error)
		Create(context.Context, resourcemodel.Comment) (*resourcemodel.Comment, error)
		CreateCopy(context.Context, resourcemodel.Comment) (resourcemodel.Comment, error)
		DeleteById(context.Context, string) error

		GetByQuery(context.Context, *resourcemodel.QueryOption) ([]*resourcemodel.Comment, int, error)
	}
)
