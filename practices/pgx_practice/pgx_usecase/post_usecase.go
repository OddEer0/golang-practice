package pgxUseCase

import (
	"context"
	"github.com/OddEer0/golang-practice/resources/sql"

	"github.com/OddEer0/golang-practice/resources/aggregate"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
)

type (
	PostUseCase interface {
		GetPostsByUserId(context.Context, domain.Id, model.PostConns) ([]aggregate.Post, error)
		GetPostById(context.Context, domain.Id, model.PostConns) (aggregate.Post, error)
	}

	postUseCase struct {
		postRepository    repository.Post
		commentRepository repository.Comment
		transactor        sql.Transactor
	}
)

// TODO - add goroutine query
func (p *postUseCase) GetPostById(ctx context.Context, id domain.Id, conns model.PostConns) (aggregate.Post, error) {
	post, err := p.postRepository.GetById(ctx, id)
	result := aggregate.Post{}

	if err != nil {
		return result, err
	}

	if opt, ok := conns["comments"]; ok {
		comments, err := p.commentRepository.GetByPostId(ctx, id, opt)
		if err != nil {
			return result, err
		}
		result.Comments = comments
	}
	result.Value = post

	return result, nil
}

// GetPostsByUserId implements PostUseCase.
func (p *postUseCase) GetPostsByUserId(context.Context, domain.Id, model.PostConns) ([]aggregate.Post, error) {
	panic("unimplemented")
}

func NewPostUseCase(postRepository repository.Post, commentRepository repository.Comment, transactor sql.Transactor) PostUseCase {
	return &postUseCase{
		postRepository:    postRepository,
		commentRepository: commentRepository,
		transactor:        transactor,
	}
}
