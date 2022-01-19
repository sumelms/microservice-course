package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Course(context.Context, uuid.UUID) (Course, error)
	Courses(context.Context) ([]Course, error)
	CreateCourse(context.Context, *Course) error
	UpdateCourse(context.Context, *Course) error
	DeleteCourse(context.Context, uuid.UUID) error
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

func (s *Service) Course(_ context.Context, id uuid.UUID) (Course, error) {
	c, err := s.repo.Course(id)
	if err != nil {
		return Course{}, fmt.Errorf("service can't find course: %w", err)
	}
	return c, nil
}

func (s *Service) Courses(_ context.Context) ([]Course, error) {
	cc, err := s.repo.Courses()
	if err != nil {
		return []Course{}, fmt.Errorf("service didn't found any course: %w", err)
	}
	return cc, nil
}

func (s *Service) CreateCourse(_ context.Context, c *Course) error {
	if err := s.repo.CreateCourse(c); err != nil {
		return fmt.Errorf("service can't create course: %w", err)
	}
	return nil
}

func (s *Service) UpdateCourse(_ context.Context, c *Course) error {
	if err := s.repo.UpdateCourse(c); err != nil {
		return fmt.Errorf("service can't update course: %w", err)
	}
	return nil
}

func (s *Service) DeleteCourse(_ context.Context, id uuid.UUID) error {
	if err := s.repo.DeleteCourse(id); err != nil {
		return fmt.Errorf("service can't delete course: %w", err)
	}
	return nil
}
