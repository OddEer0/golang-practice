package pgxRepository

import (
	"context"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/OddEer0/golang-practice/resources/sql"
	"log/slog"
	"strings"
	"time"
)

type userRepository struct {
	txController sql.TransactionController
	db           sql.QueryExecutor
}

func (u *userRepository) GetByQuery(ctx context.Context, opt *model.ManyOpt) ([]*model.User, error) {
	db := u.getQueryExecutor(ctx)
	limit := opt.Limit
	page := (opt.Page - 1) * opt.Limit
	queryStr := strings.Builder{}
	queryStr.WriteString(GetUserByQueryPart1)
	queryStr.WriteString(opt.SortBy)
	queryStr.WriteString(" ")
	queryStr.WriteString(opt.SortDir)
	queryStr.WriteString(GetUserByQueryPart12)

	rows, err := db.Query(ctx, queryStr.String(), limit, page)
	if err != nil {
		slog.Error("GetUserByQuery err", slog.Any("cause", err))
		return nil, domain.ErrInternal
	}
	defer rows.Close()
	users := make([]*model.User, 0, opt.Limit)
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(
			&user.Id,
			&user.Login,
			&user.Email,
			&user.Password,
			&user.UpdatedAt,
			&user.CreatedAt,
		)
		if err != nil {
			slog.Error("GetUserByQuery rows.Scan err", slog.Any("cause", err))
			return nil, domain.ErrInternal
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		slog.Error("GetUserByQuery rows.Err", slog.Any("cause", err))
		return nil, domain.ErrInternal
	}
	return users, nil
}

func (u *userRepository) getQueryExecutor(ctx context.Context) sql.QueryExecutor {
	query := u.txController.ExtractTransaction(ctx)
	if query != nil {
		return query
	}
	return u.db
}

func (u *userRepository) GetById(ctx context.Context, id domain.Id) (*model.User, error) {
	db := u.getQueryExecutor(ctx)
	user := &model.User{}
	err := db.QueryRow(ctx, GetUserById, id).Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Password,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, domain.ErrInternal
	}
	return user, nil
}

func (u *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	db := u.getQueryExecutor(ctx)
	_, err := db.Exec(ctx, CreateUserQuery, user.Id, user.Login, user.Email, user.Password, user.UpdatedAt, user.CreatedAt)
	if err != nil {
		slog.Error("create user error", slog.Any("cause", err))
		return nil, domain.ErrInternal
	}
	return user, nil
}

func (u *userRepository) UpdateById(ctx context.Context, user *model.User) (*model.User, error) {
	db := u.getQueryExecutor(ctx)
	updatedTime := time.Now()
	_, err := db.Exec(ctx, UpdateUserById, user.Login, user.Email, user.Password, updatedTime, user.Id)
	if err != nil {
		return nil, domain.ErrInternal
	}
	user.UpdatedAt = updatedTime
	return user, nil
}

func (u *userRepository) UpdateLoginById(ctx context.Context, id domain.Id, login string) (*model.User, error) {
	db := u.getQueryExecutor(ctx)
	user := &model.User{}
	err := db.QueryRow(ctx, UpdateUserLoginById, login, id).Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Password,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, domain.ErrInternal
	}
	return user, nil
}

func (u *userRepository) DeleteById(ctx context.Context, id domain.Id) error {
	db := u.getQueryExecutor(ctx)
	_, err := db.Exec(ctx, DeleteUserById, id)
	if err != nil {
		return domain.ErrInternal
	}
	return nil
}

func NewUserRepository(db sql.QueryExecutor, txController sql.TransactionController) repository.User {
	return &userRepository{
		txController: txController,
		db:           db,
	}
}
