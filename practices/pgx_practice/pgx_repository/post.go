package pgxRepository

import (
	"context"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/OddEer0/golang-practice/resources/sql"
	"github.com/google/uuid"
	"time"
)

type postRepository struct {
	txController sql.TransactionController
	db           sql.QueryExecutor
}

func (p *postRepository) getQueryExecutor(ctx context.Context) sql.QueryExecutor {
	query := p.txController.ExtractTransaction(ctx)
	if query != nil {
		return query
	}
	return p.db
}

func (p *postRepository) GetById(ctx context.Context, id domain.Id) (*model.Post, error) {
	db := p.getQueryExecutor(ctx)
	post := &model.Post{}
	err := db.QueryRow(ctx, GetPostByIdQuery, id).Scan(
		&post.Id,
		&post.OwnerId,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return nil, domain.ErrInternal
	}
	return post, nil
}

func (p *postRepository) GetByOwnerId(ctx context.Context, id domain.Id, opt *model.ManyOpt) ([]*model.Post, error) {
	db := p.getQueryExecutor(ctx)
	limit := opt.Limit
	page := opt.Page * (opt.Limit + 1)
	queryStr := GetPostByQueryAsc
	if opt.SortDir == "Desc" {
		queryStr += GetPostByQueryDesc
	}
	rows, err := db.Query(ctx, queryStr, id, opt.SortBy, limit, page)
	if err != nil {
		return nil, domain.ErrInternal
	}
	defer rows.Close()
	var posts []*model.Post
	for rows.Next() {
		post := &model.Post{}
		err := rows.Scan(
			&post.Id,
			&post.OwnerId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, domain.ErrInternal
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, domain.ErrInternal
	}

	return posts, nil
}

func (p *postRepository) Create(ctx context.Context, post *model.Post) (*model.Post, error) {
	db := p.getQueryExecutor(ctx)
	_, err := db.Exec(ctx, CreatePostQuery, post.Id, post.OwnerId, post.Title, post.Content, post.UpdatedAt, post.CreatedAt)
	if err != nil {
		return nil, domain.ErrInternal
	}
	return post, nil
}

func (p *postRepository) UpdateById(ctx context.Context, post *model.Post) (*model.Post, error) {
	db := p.getQueryExecutor(ctx)
	updatedTime := time.Now()
	_, err := db.Exec(ctx, UpdatePostById, post.Title, post.Content, updatedTime, post.Id)
	if err != nil {
		return nil, domain.ErrInternal
	}
	post.UpdatedAt = updatedTime
	return post, nil
}

func (p *postRepository) DeleteById(ctx context.Context, id domain.Id) error {
	db := p.getQueryExecutor(ctx)
	_, err := db.Exec(ctx, DeletePostByIdQuery, id)
	if err != nil {
		return domain.ErrInternal
	}
	return nil
}

func (p *postRepository) UpdateBodyByUserId(ctx context.Context, user *model.User) (*model.Post, error) {
	uniq := uuid.New().String()
	db := p.getQueryExecutor(ctx)
	post := &model.Post{}
	err := db.QueryRow(ctx, `UPDATE posts SET title = $1 WHERE ownerId = $2 RETURNING id, ownerId, title, content, updatedAt, createdAt`, uniq, user.Id).Scan(
		&post.Id,
		&post.OwnerId,
		&post.Title,
		&post.Content,
		&post.UpdatedAt,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, domain.ErrInternal
	}
	return post, nil
}

func NewPostRepository(db sql.QueryExecutor, txController sql.TransactionController) repository.Post {
	return &postRepository{
		txController: txController,
		db:           db,
	}
}
