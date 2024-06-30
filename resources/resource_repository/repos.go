package resourceRepository

import (
	"context"

	resourceModel "github.com/OddEer0/golang-practice/resources/resource_model"
)

type (
	User interface {
		GetById(context.Context, string) (*resourceModel.User, error)
		GetByIdCopy(context.Context, string) (resourceModel.User, error)
		Create(context.Context, resourceModel.User) (*resourceModel.User, error)
		CreateCopy(context.Context, resourceModel.User) (resourceModel.User, error)
		DeleteById(context.Context, string) error
	}
	Post interface {
		GetById(context.Context, string) (*resourceModel.Post, error)
		GetByIdCopy(context.Context, string) (resourceModel.Post, error)
		Create(context.Context, resourceModel.Post) (*resourceModel.Post, error)
		CreateCopy(context.Context, resourceModel.Post) (resourceModel.Post, error)
		DeleteById(context.Context, string) error
	}
	Comment interface {
		GetById(context.Context, string) (*resourceModel.Comment, error)
		GetByIdCopy(context.Context, string) (resourceModel.Comment, error)
		Create(context.Context, resourceModel.Comment) (*resourceModel.Comment, error)
		CreateCopy(context.Context, resourceModel.Comment) (resourceModel.Comment, error)
		DeleteById(context.Context, string) error

		GetByQuery(context.Context, *resourceModel.QueryOption) ([]*resourceModel.Comment, int, error)
	}
)
