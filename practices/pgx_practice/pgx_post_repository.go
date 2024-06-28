package pgxpractice

import (
	"context"

	resourcemodel "github.com/OddEer0/golang-practice/resources/resource_model"
	resourcerepository "github.com/OddEer0/golang-practice/resources/resource_repository"
)

type postRepository struct {
	db Query
}

// Create implements resourcerepository.PostRepository.
func (p postRepository) Create(context.Context, resourcemodel.Post) (*resourcemodel.Post, error) {
	panic("unimplemented")
}

// CreateCopy implements resourcerepository.PostRepository.
func (p postRepository) CreateCopy(context.Context, resourcemodel.Post) (resourcemodel.Post, error) {
	panic("unimplemented")
}

// DeleteById implements resourcerepository.PostRepository.
func (p postRepository) DeleteById(context.Context, string) error {
	panic("unimplemented")
}

// GetById implements resourcerepository.PostRepository.
func (p postRepository) GetById(context.Context, string) (*resourcemodel.Post, error) {
	panic("unimplemented")
}

// GetByIdCopy implements resourcerepository.PostRepository.
func (p postRepository) GetByIdCopy(context.Context, string) (resourcemodel.Post, error) {
	panic("unimplemented")
}

// GetPostAggregate implements resourcerepository.PostRepository.
func (p postRepository) GetPostAggregate(ctx context.Context, id string, opt *resourcemodel.AggregateOption) (*resourcemodel.PostAggregate, error) {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

}

// GetPostAggregateCopy implements resourcerepository.PostRepository.
func (p postRepository) GetPostAggregateCopy(context.Context, string, *resourcemodel.AggregateOption) (resourcemodel.PostAggregate, error) {
	panic("unimplemented")
}

func NewPostRepository(conn Query) resourcerepository.PostRepository {
	return postRepository{
		db: conn,
	}
}
