package pgxInteractor

import "github.com/OddEer0/golang-practice/resources/repository"

type (
	Interactor struct {
		userRepository repository.User
		postRepository repository.Post
		commentRepository repository.Comment
	}
)

