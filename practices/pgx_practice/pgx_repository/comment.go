package pgxRepository

import (
	"context"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/OddEer0/golang-practice/resources/sql"
)

type commentRepository struct {
	db           sql.QueryExecutor
	txController sql.TransactionController
}

func (c *commentRepository) GetById(ctx context.Context, id domain.Id) (*model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepository) GetByPostId(ctx context.Context, id domain.Id, opt *model.ManyOpt) ([]*model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepository) Create(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepository) UpdateById(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepository) DeleteById(ctx context.Context, id domain.Id) error {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepository) UpdateBodyByUserId(ctx context.Context, user *model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewCommentRepository(db sql.QueryExecutor, txController sql.TransactionController) repository.Comment {
	return &commentRepository{}
}
