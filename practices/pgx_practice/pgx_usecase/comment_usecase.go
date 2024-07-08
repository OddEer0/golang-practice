package pgxUseCase

import (
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/OddEer0/golang-practice/resources/sql"
)

type commentUseCase struct {
	userRepository    repository.User
	postRepository    repository.Post
	commentRepository repository.Comment
	transactor        sql.Transactor
}

type CommentUseCase interface{}

func NewCommentUseCase(commentRepository repository.Comment, postRepository repository.Post, userRepository repository.User, transactor sql.Transactor) CommentUseCase {
	return &commentUseCase{
		userRepository:    userRepository,
		postRepository:    postRepository,
		commentRepository: commentRepository,
		transactor:        transactor,
	}
}
