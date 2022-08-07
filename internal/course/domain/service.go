package domain

import (
	"context"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

// Service defines the domains service interface
type Service interface {
	Course(ctx context.Context, courseID uuid.UUID) (Course, error)
	Courses(ctx context.Context) ([]Course, error)
	CreateCourse(ctx context.Context, c *Course) error
	UpdateCourse(ctx context.Context, c *Course) error
	DeleteCourse(ctx context.Context, courseID uuid.UUID) error

	SubscribeCourse(ctx context.Context, cs *Subscription) error
	UnsubscribeCourse(ctx context.Context, courseID uuid.UUID, userID uuid.UUID) error
}

type serviceConfiguration func(svc *service) error

type service struct {
	courses       CourseRepository
	subscriptions SubscriptionRepository
	logger        log.Logger
}

// NewService creates a new domain service instance
func NewService(cfgs ...serviceConfiguration) (*service, error) {
	svc := &service{}
	for _, cfg := range cfgs {
		err := cfg(svc)
		if err != nil {
			return nil, err
		}
	}
	return svc, nil
}

// WithCourseRepository injects the course repository to the domain service
func WithCourseRepository(cr CourseRepository) serviceConfiguration {
	return func(svc *service) error {
		svc.courses = cr
		return nil
	}
}

// WithSubscriptionRepository injects the subscription repository to the domain service
func WithSubscriptionRepository(sr SubscriptionRepository) serviceConfiguration {
	return func(svc *service) error {
		svc.subscriptions = sr
		return nil
	}
}

// WithLogger injects the logger to the domain service
func WithLogger(l log.Logger) serviceConfiguration {
	return func(svc *service) error {
		svc.logger = l
		return nil
	}
}
