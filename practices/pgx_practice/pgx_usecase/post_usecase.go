package pgxUseCase

import (
	"context"
	"github.com/OddEer0/golang-practice/resources/aggregate"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/OddEer0/golang-practice/resources/sql"
	"log/slog"
)

type (
	PostUseCase interface {
		GetPostsByUserId(context.Context, domain.Id, *model.ManyOpt, model.PostConns) ([]*aggregate.Post, error)
		GetPostById(context.Context, domain.Id, model.PostConns) (aggregate.Post, error)
	}

	postUseCase struct {
		postRepository    repository.Post
		commentRepository repository.Comment
		userRepository    repository.User
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

func (p *postUseCase) GetPostsByUserId(ctx context.Context, id domain.Id, opt *model.ManyOpt, conns model.PostConns) ([]*aggregate.Post, error) {
	posts, err := p.postRepository.GetByOwnerId(ctx, id, opt)
	if err != nil {
		return nil, err
	}
	postAggregates := make([]*aggregate.Post, 0, len(posts))
	for _, post := range posts {
		postAggregate := &aggregate.Post{
			Value: post,
		}

		if _, ok := conns["owner"]; ok {
			owner, err := p.userRepository.GetById(ctx, post.OwnerId)
			if err != nil {
				return nil, err
			}
			postAggregate.Owner = owner
		}

		if option, ok := conns["comments"]; ok {
			slog.Info("comment option", slog.Any("option", option))
			comments, err := p.commentRepository.GetByPostId(ctx, post.Id, option)
			slog.Info("comments by post id", slog.Any("comments", comments))
			if err != nil {
				return nil, err
			}
			postAggregate.Comments = comments
		}

		postAggregates = append(postAggregates, postAggregate)
	}
	return postAggregates, nil
}

func NewPostUseCase(postRepository repository.Post, commentRepository repository.Comment, transactor sql.Transactor) PostUseCase {
	return &postUseCase{
		postRepository:    postRepository,
		commentRepository: commentRepository,
		transactor:        transactor,
	}
}
