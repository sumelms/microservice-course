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
	CreateCourse(ctx context.Context, c *Course) error
	UpdateCourse(ctx context.Context, c *Course) error
	DeleteCourse(ctx context.Context, courseID uuid.UUID) error

	Subscription(ctx context.Context, id uuid.UUID) (Subscription, error)
	Subscriptions(ctx context.Context) ([]Subscription, error)
	CreateSubscription(ctx context.Context, cs *Subscription) error
	UpdateSubscription(ctx context.Context, cs *Subscription) error
	DeleteSubscription(ctx context.Context, id uuid.UUID) error
}

type ServiceConfiguration func(svc *Service) error

type Service struct {
	courses       CourseRepository
	subscriptions SubscriptionRepository
	logger        log.Logger
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

// WithSubscriptionRepository injects the subscription repository to the domain Service.
func WithSubscriptionRepository(sr SubscriptionRepository) ServiceConfiguration {
	return func(svc *Service) error {
		svc.subscriptions = sr

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
