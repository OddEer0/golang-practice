package main

import (
	"context"
	"fmt"

	pgxPractice "github.com/OddEer0/golang-practice/practices/pgx_practice"
)

func main() {
	cfg := pgxPractice.NewConfig()
	connPool := pgxPractice.ConnectPg(cfg)

	ctx := context.Background()

	// id := uuid.New().String()

	// _, err := connPool.Exec(ctx, `
	// 	INSERT INTO users (id, login, email, password, updatedAt, createdAt)
	// 	VALUES ($1, $2, $3, $4, $5, $6);
	// `, id, "eer0", "eer0@gmail.com", "kekw1234", time.Now(), time.Now())

	// if err != nil {
	// 	panic(err)
	// }

	var login string
	err := connPool.QueryRow(ctx, `
		SELECT login FROM users
		WHERE login = $1;
	`, "eer0").Scan(&login)
	if err != nil {
		panic(err)
	}
	fmt.Println("get user login:", login)
}
