package pgxUseCase

import (
	"context"
	"github.com/OddEer0/golang-practice/resources/aggregate"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/OddEer0/golang-practice/resources/sql"
)

type commentUseCase struct {
	userRepository    repository.User
	postRepository    repository.Post
	commentRepository repository.Comment
	transactor        sql.Transactor
}

func (c *commentUseCase) GetCommentsByOwnerId(ctx context.Context, postId domain.Id, opt *model.ManyOpt, connOpt model.CommentConns) ([]*aggregate.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentUseCase) GetCommentsByPostId(ctx context.Context, postId domain.Id, opt *model.ManyOpt, connOpt model.CommentConns) ([]*aggregate.Comment, error) {
	comments, err := c.commentRepository.GetByPostId(ctx, postId, opt)
	if err != nil {
		return nil, err
	}
	commentAggregates := make([]*aggregate.Comment, 0, len(comments))
	for _, comment := range comments {
		commentAggregate := aggregate.Comment{
			Value: comment,
		}
		if _, ok := connOpt["owner"]; ok {
			owner, err := c.userRepository.GetById(ctx, comment.OwnerId)
			if err != nil {
				return nil, err
			}
			commentAggregate.Owner = owner
		}
		if _, ok := connOpt["post"]; ok {
			post, err := c.postRepository.GetById(ctx, comment.PostId)
			if err != nil {
				return nil, err
			}
			commentAggregate.Post = post
		}
		commentAggregates = append(commentAggregates, &commentAggregate)
	}
	return commentAggregates, nil
}

type CommentUseCase interface {
	GetCommentsByPostId(ctx context.Context, postId domain.Id, opt *model.ManyOpt, connOpt model.CommentConns) ([]*aggregate.Comment, error)
	GetCommentsByOwnerId(ctx context.Context, postId domain.Id, opt *model.ManyOpt, connOpt model.CommentConns) ([]*aggregate.Comment, error)
}

func NewCommentUseCase(commentRepository repository.Comment, postRepository repository.Post, userRepository repository.User, transactor sql.Transactor) CommentUseCase {
	return &commentUseCase{
		userRepository:    userRepository,
		postRepository:    postRepository,
		commentRepository: commentRepository,
		transactor:        transactor,
	}
}
