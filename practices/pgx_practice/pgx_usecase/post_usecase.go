package pgxUseCase

import (
	"context"

	"github.com/OddEer0/golang-practice/resources/aggregate"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/repository"
)

type (
	PostUseCase interface {
		GetPostsByUserId(context.Context, domain.Id, aggregate.PostConns) ([]aggregate.Post, error)
		GetPostById(context.Context, domain.Id, aggregate.PostConns) (aggregate.Post, error)
	}

	postUseCase struct {
		postRepository    repository.Post
		commentRepository repository.Comment
	}
)

// GetPostById implements PostUseCase.
func (p *postUseCase) GetPostById(context.Context, domain.Id, aggregate.PostConns) (aggregate.Post, error) {
	panic("unimplemented")
}

// GetPostsByUserId implements PostUseCase.
func (p *postUseCase) GetPostsByUserId(context.Context, domain.Id, aggregate.PostConns) ([]aggregate.Post, error) {
	panic("unimplemented")
}

func NewPostUseCase(postRepository repository.Post, commentRepository repository.Comment) PostUseCase {
	return &postUseCase{
		postRepository:    postRepository,
		commentRepository: commentRepository,
	}
}
