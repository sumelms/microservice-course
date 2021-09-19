package domain

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
)

type Service interface {
	CreateCourse(context.Context, *Course) (Course, error)
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

func (s *service) CreateCourse(_ context.Context, course *Course) (Course, error) {
	p, err := s.repo.Create(course)
	if err != nil {
		return Course{}, fmt.Errorf("create course: %w", err)
	}
	return p, nil
}
