package runners

import (
	"context"

	pgxPractice "github.com/OddEer0/golang-practice/practices/pgx_practice"
)

const (
	CreateUser = `
		INSERT INTO users (id, login, password)
		VALUES ($1, $2, $3)
		RETURNING id, login;
	`
	UpdateUserLogin = `
		UPDATE users SET login = $2
		WHERE id = $1
		RETURNING id, login;
	`
	UpdatePostContentById = `
		UPDATE posts SET title = $2, content = $3
		WHERE ownerId = $1
	`
	GetUserPostsId = `
		UPDATE id FROM posts
		WHERE ownerId = $1;
	`
)

func RunPgx() {
	connPool := pgxPractice.ConnectPg("")
	ctx := context.Background()

	connPool.QueryRow(ctx, CreateUser, "one", "eer0", "12345678")
}
