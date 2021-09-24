package domain

import (
	"context"

	"github.com/go-kit/kit/log"
)

type Service interface {
	ListMatrix(context.Context) ([]Matrix, error)
	CreateMatrix(context.Context, *Matrix) (Matrix, error)
	FindMatrix(context.Context, string) (Matrix, error)
	UpdateMatrix(context.Context, *Matrix) (Matrix, error)
	DeleteMatrix(context.Context, string) error
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) *service { // nolint: revive
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s service) ListMatrix(ctx context.Context) ([]Matrix, error) {
	panic("implement me")
}

func (s service) CreateMatrix(ctx context.Context, matrix *Matrix) (Matrix, error) {
	panic("implement me")
}

func (s service) FindMatrix(ctx context.Context, id string) (Matrix, error) {
	panic("implement me")
}

func (s service) UpdateMatrix(ctx context.Context, matrix *Matrix) (Matrix, error) {
	panic("implement me")
}

func (s service) DeleteMatrix(ctx context.Context, id string) error {
	panic("implement me")
}
