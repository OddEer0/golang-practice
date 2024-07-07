package model

import (
	"time"

	"github.com/OddEer0/golang-practice/resources/domain"
)

type (
	Id string

	User struct {
		Id domain.Id
		Login string
		Password string
		Email string
		UpdatedAt time.Time
		CreatedAt time.Time
	}

	Post struct {
		Id domain.Id
		OwnerId domain.Id
		Title string
		Content string
		UpdatedAt time.Time
		CreatedAt time.Time
	}

	Comment struct {
		Id domain.Id
		OwnerId domain.Id
		PostId domain.Id
		Message string
		UpdatedAt time.Time
		CreatedAt time.Time
	}
)
