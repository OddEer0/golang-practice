package pgxPractice

import (
	"context"

	resourceModel "github.com/OddEer0/golang-practice/resources/resource_model"
	resourceRepository "github.com/OddEer0/golang-practice/resources/resource_repository"
)

type PgxUserUseCase interface {
	UpdateUserLogin(ctx context.Context, id, newLogin string) (*resourceModel.User, error)
}

type pgxUserUseCase struct {
	userRepository resourceRepository.User
	postRepository resourceRepository.Post
}

// UpdateUserLogin implements PgxUserUseCase.
func (p *pgxUserUseCase) UpdateUserLogin(ctx context.Context, id string, newLogin string) (*resourceModel.User, error) {
	panic("kek")	
}

func NewPgxUserUseCase(postRepository resourceRepository.Post, userRepository resourceRepository.User) PgxUserUseCase {
	return &pgxUserUseCase{
		postRepository: postRepository,
		userRepository: userRepository,
	}
}
