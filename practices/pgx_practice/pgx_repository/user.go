package pgxRepository

import (
	"context"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/OddEer0/golang-practice/resources/sql"
)

type userRepository struct {
	txController sql.TransactionController
	db           sql.QueryExecutor
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
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	db := u.getQueryExecutor(ctx)
	_, err := db.Exec(ctx, CreateUserQuery, user.Id, user.Login, user.Email, user.Password, user.UpdatedAt, user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) UpdateById(ctx context.Context, user *model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
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
		return nil, err
	}
	return user, nil
}

func (u *userRepository) DeleteById(ctx context.Context, id domain.Id) error {
	db := u.getQueryExecutor(ctx)
	var err error
	_, err = db.Exec(ctx, DeleteUserById, id)
	return err
}

func NewUserRepository(db sql.QueryExecutor, txController sql.TransactionController) repository.User {
	return &userRepository{
		txController: txController,
		db:           db,
	}
}
