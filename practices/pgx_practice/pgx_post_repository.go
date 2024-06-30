package pgxPractice

import (
	"context"

	resourceModel "github.com/OddEer0/golang-practice/resources/resource_model"
	resourceRepository "github.com/OddEer0/golang-practice/resources/resource_repository"
)

type postRepository struct {
	db QueryExecutor
}

// Create implements resourceRepository.PostRepository.
func (p postRepository) Create(context.Context, resourceModel.Post) (*resourceModel.Post, error) {
	panic("unimplemented")
}

// CreateCopy implements resourceRepository.PostRepository.
func (p postRepository) CreateCopy(context.Context, resourceModel.Post) (resourceModel.Post, error) {
	panic("unimplemented")
}

// DeleteById implements resourceRepository.PostRepository.
func (p postRepository) DeleteById(context.Context, string) error {
	panic("unimplemented")
}

// GetById implements resourceRepository.PostRepository.
func (p postRepository) GetById(context.Context, string) (*resourceModel.Post, error) {
	panic("unimplemented")
}

// GetByIdCopy implements resourceRepository.PostRepository.
func (p postRepository) GetByIdCopy(context.Context, string) (resourceModel.Post, error) {
	panic("unimplemented")
}

func NewPostRepository(conn QueryExecutor) resourceRepository.Post {
	return postRepository{
		db: conn,
	}
}
