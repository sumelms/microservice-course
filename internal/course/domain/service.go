package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type ServiceInterface interface {
	ListCourse(context.Context) ([]Course, error)
	CreateCourse(context.Context, *Course) (Course, error)
	FindCourse(context.Context, string) (Course, error)
	UpdateCourse(context.Context, *Course) (Course, error)
	DeleteCourse(context.Context, string) error
}

type Service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) ListCourse(_ context.Context) ([]Course, error) {
	cs, err := s.repo.List()
	if err != nil {
		return []Course{}, fmt.Errorf("Service didn't found any course: %w", err)
	}
	return cs, nil
}

func (s *Service) CreateCourse(_ context.Context, course *Course) (Course, error) {
	c, err := s.repo.Create(course)
	if err != nil {
		return Course{}, fmt.Errorf("Service can't create course: %w", err)
	}
	return c, nil
}

func (s *Service) FindCourse(_ context.Context, id string) (Course, error) {
	c, err := s.repo.Find(id)
	if err != nil {
		return Course{}, fmt.Errorf("Service can't find course: %w", err)
	}
	return c, nil
}

func (s *Service) UpdateCourse(_ context.Context, course *Course) (Course, error) {
	c, err := s.repo.Update(course)
	if err != nil {
		return Course{}, fmt.Errorf("Service can't update course: %w", err)
	}
	return c, nil
}

func (s *Service) DeleteCourse(_ context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("Service can't delete course: %w", err)
	}
	return nil
}
