package domain

import (
	"context"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

// ServiceInterface defines the domains Service interface.
type ServiceInterface interface {
	Course(ctx context.Context, id uuid.UUID) (Course, error)
	Courses(ctx context.Context) ([]Course, error)
	CreateCourse(ctx context.Context, course *Course) error
	UpdateCourse(ctx context.Context, course *Course) error
	DeleteCourse(ctx context.Context, course *DeletedCourse) error
}

type ServiceConfiguration func(svc *Service) error

type Service struct {
	courses CourseRepository
	logger  log.Logger
}

// NewService creates a new domain Service instance.
func NewService(cfgs ...ServiceConfiguration) (*Service, error) {
	svc := &Service{}
	for _, cfg := range cfgs {
		err := cfg(svc)
		if err != nil {
			return nil, err
		}
	}

	return svc, nil
}

// WithCourseRepository injects the course repository to the domain Service.
func WithCourseRepository(cr CourseRepository) ServiceConfiguration {
	return func(svc *Service) error {
		svc.courses = cr

		return nil
	}
}

// WithLogger injects the logger to the domain Service.
func WithLogger(l log.Logger) ServiceConfiguration {
	return func(svc *Service) error {
		svc.logger = l

		return nil
	}
}
