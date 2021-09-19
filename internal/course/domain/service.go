package domain

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
)

type Service interface {
	CreateCourse(context.Context, *Course) (Course, error)
	GetCourse(context.Context, string) (Course, error)
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
	c, err := s.repo.Create(course)
	if err != nil {
		return Course{}, fmt.Errorf("service can't create course: %w", err)
	}
	return c, nil
}

func (s *service) GetCourse(_ context.Context, id string) (Course, error) {
	c, err := s.repo.Find(id)
	if err != nil {
		return Course{}, fmt.Errorf("service can't found course: %w", err)
	}
	return c, nil
}
