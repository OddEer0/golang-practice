package pgxRepository

import (
	"context"
	"fmt"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/OddEer0/golang-practice/resources/sql"
	"log/slog"
	"strings"
	"time"
)

type commentContent struct {
	Id      domain.Id
	Content string
}

type commentRepository struct {
	db           sql.QueryExecutor
	txController sql.TransactionController
}

func (c *commentRepository) GetByOwnerId(ctx context.Context, id domain.Id, opt *model.ManyOpt) ([]*model.Comment, error) {
	db := c.getQueryExecutor(ctx)
	page := (opt.Page - 1) * opt.Limit
	queryStr := strings.Builder{}
	queryStr.WriteString(GetCommentsByOwnerIdQueryPart1)
	queryStr.WriteString(opt.SortBy)
	queryStr.WriteString(" ")
	queryStr.WriteString(opt.SortDir)
	queryStr.WriteString(GetCommentsByOwnerIdQueryPart2)
	rows, err := db.Query(ctx, queryStr.String(), id, opt.SortBy, opt.Limit, page)
	if err != nil {
		return nil, domain.ErrInternal
	}
	defer rows.Close()
	comments := make([]*model.Comment, 0, opt.Limit)
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
	page := (opt.Page - 1) * opt.Limit
	queryStr := strings.Builder{}
	queryStr.WriteString(GetCommentsByPostIdQueryPart1)
	queryStr.WriteString(opt.SortBy)
	queryStr.WriteString(" ")
	queryStr.WriteString(opt.SortDir)
	queryStr.WriteString(GetCommentsByPostIdQueryPart2)
	rows, err := db.Query(ctx, queryStr.String(), id, opt.Limit, page)
	if err != nil {
		slog.Error("query comment by post id err", slog.Any("cause", err))
		return nil, domain.ErrInternal
	}
	defer rows.Close()
	comments := make([]*model.Comment, 0, opt.Limit)
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
			slog.Error("scan fail", slog.Any("cause", err))
			return nil, domain.ErrInternal
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		slog.Error("rows err fail", slog.Any("cause", err))
		return nil, domain.ErrInternal
	}
	return comments, nil
}

func (c *commentRepository) Create(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	db := c.getQueryExecutor(ctx)
	_, err := db.Exec(ctx, CreateCommentQuery, comment.Id, comment.OwnerId, comment.PostId, comment.Message, comment.UpdatedAt, comment.CreatedAt)
	if err != nil {
		slog.Error("create comment error", slog.Any("cause", err))
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
	slog.Info("update comment by user id ", slog.Any("user", user))
	db := c.getQueryExecutor(ctx)
	rows, err := db.Query(ctx, GetAllCommentsIdContentByUserId, user.Id)
	if err != nil {
		slog.Error("query comment by user id err", slog.Any("cause", err))
		return domain.ErrInternal
	}
	defer rows.Close()
	comm := make([]*commentContent, 0, 50)
	for rows.Next() {
		comment := &commentContent{}
		err := rows.Scan(&comment.Id, &comment.Content)
		if err != nil {
			slog.Error("scan fail", slog.Any("cause", err))
			return domain.ErrInternal
		}
		comm = append(comm, comment)
	}

	fmt.Println(comm)

	updatedTime := time.Now()
	for _, comment := range comm {
		word := strings.Split(comment.Content, " ")
		newWords := make([]string, 0, len(word))
		for _, word := range word {
			if strings.HasPrefix(word, "<login>") && strings.HasSuffix(word, "<login>") {
				newWords = append(newWords, "<login>"+user.Login+"<login>")
			} else {
				newWords = append(newWords, word)
			}
		}
		_, err := db.Exec(ctx, UpdateCommentByIdQuery, strings.Join(newWords, " "), updatedTime, comment.Id)
		if err != nil {
			return domain.ErrInternal
		}
	}

	return nil
}

func NewCommentRepository(db sql.QueryExecutor, txController sql.TransactionController) repository.Comment {
	return &commentRepository{
		db:           db,
		txController: txController,
	}
}
