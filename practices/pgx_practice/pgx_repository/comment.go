package pgxRepository

import (
	"context"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/OddEer0/golang-practice/resources/sql"
	"strings"
	"time"
)

type commentContent struct {
	id      domain.Id
	content string
}

type commentRepository struct {
	db           sql.QueryExecutor
	txController sql.TransactionController
}

func (c *commentRepository) getQueryExecutor(ctx context.Context) sql.QueryExecutor {
	query := c.txController.ExtractTransaction(ctx)
	if query != nil {
		return query
	}
	return c.db
}

func (c *commentRepository) GetById(ctx context.Context, id domain.Id) (*model.Comment, error) {
	comment := &model.Comment{}
	db := c.getQueryExecutor(ctx)
	err := db.QueryRow(ctx, GetCommentByIdQuery, id).Scan(
		&comment.Id,
		&comment.OwnerId,
		&comment.PostId,
		&comment.Message,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		return nil, domain.ErrInternal
	}
	return comment, nil
}

func (c *commentRepository) GetByPostId(ctx context.Context, id domain.Id, opt *model.ManyOpt) ([]*model.Comment, error) {
	db := c.getQueryExecutor(ctx)
	queryStr := GetCommentsByPostIdQueryAsc
	if opt.SortDir == "Desc" {
		queryStr = GetCommentsByPostIdQueryDesc
	}
	page := opt.Page * (opt.Limit + 1)
	rows, err := db.Query(ctx, queryStr, id, opt.SortBy, opt.Limit, page)
	if err != nil {
		return nil, domain.ErrInternal
	}
	comments := make([]*model.Comment, 0, opt.Limit)
	defer rows.Close()
	for rows.Next() {
		comment := &model.Comment{}
		err := rows.Scan(
			&comment.Id,
			&comment.OwnerId,
			&comment.PostId,
			&comment.Message,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, domain.ErrInternal
		}
	}
	if err := rows.Err(); err != nil {
		return nil, domain.ErrInternal
	}
	return comments, nil
}

func (c *commentRepository) Create(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	db := c.getQueryExecutor(ctx)
	_, err := db.Exec(ctx, CreateCommentQuery, comment.Id, comment.OwnerId, comment.PostId, comment.Message, comment.UpdatedAt, comment.CreatedAt)
	if err != nil {
		return nil, domain.ErrInternal
	}
	return comment, nil
}

func (c *commentRepository) UpdateById(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	db := c.getQueryExecutor(ctx)
	updatedTime := time.Now()
	_, err := db.Exec(ctx, UpdateCommentByIdQuery, comment.Message, updatedTime, comment.Id)
	if err != nil {
		return nil, domain.ErrInternal
	}
	comment.UpdatedAt = updatedTime
	return comment, nil
}

func (c *commentRepository) DeleteById(ctx context.Context, id domain.Id) error {
	db := c.getQueryExecutor(ctx)
	_, err := db.Exec(ctx, DeleteCommentByIdQuery, id)
	if err != nil {
		return domain.ErrInternal
	}
	return nil
}

func (c *commentRepository) UpdateBodyByUserId(ctx context.Context, user *model.User) error {
	db := c.getQueryExecutor(ctx)
	rows, err := db.Query(ctx, GetAllCommentsIdContentByPostId, user.Id)
	if err != nil {
		return domain.ErrInternal
	}
	defer rows.Close()
	comm := make([]*commentContent, 0, 50)
	for rows.Next() {
		comment := &commentContent{}
		err := rows.Scan(&comment.id, &comment.content)
		if err != nil {
			return domain.ErrInternal
		}
		comm = append(comm, comment)
	}

	updatedTime := time.Now()
	for _, comment := range comm {
		word := strings.Split(comment.content, " ")
		newWords := make([]string, 0, len(word))
		for _, word := range word {
			if strings.HasPrefix(word, "<login>") && strings.HasSuffix(word, "<login>") {
				newWords = append(newWords, "<login>"+user.Login+"<login>")
			}
			newWords = append(newWords, word)
		}
		_, err := db.Exec(ctx, UpdateCommentByIdQuery, comment.content, updatedTime, comment.id)
		if err != nil {
			return domain.ErrInternal
		}
	}

	return nil
}

func NewCommentRepository(db sql.QueryExecutor, txController sql.TransactionController) repository.Comment {
	return &commentRepository{}
}
