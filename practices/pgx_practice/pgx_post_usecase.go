package pgxPractice

import (
	"context"

	resourceModel "github.com/OddEer0/golang-practice/resources/resource_model"
	resourcePort "github.com/OddEer0/golang-practice/resources/resource_port"
	resourceRepository "github.com/OddEer0/golang-practice/resources/resource_repository"
)

type PgxPostUseCase interface {
	GetPostAggregate(ctx context.Context, id string, opt *resourceModel.AggregateOption) (*resourceModel.PostAggregate, error)
}

type pgxPostUseCase struct {
	postRepository resourceRepository.Post
	transactor     resourcePort.Transactor
}

// GetPostAggregate implements PgxPostUseCase.
func (p *pgxPostUseCase) GetPostAggregate(ctx context.Context, id string, opt *resourceModel.AggregateOption) (*resourceModel.PostAggregate, error) {
	err := p.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		return nil
	})

	return nil, err
}

func NewPgxPostUseCase(postRepository resourceRepository.Post, transactor resourcePort.Transactor) PgxPostUseCase {
	return &pgxPostUseCase{
		postRepository: postRepository,
		transactor:     transactor,
	}
}
