package domain

import (
	"context"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

// ServiceInterface defines the domains Service interface.
type ServiceInterface interface {
	Subscription(ctx context.Context, subscriptionUUID uuid.UUID) (Subscription, error)
	Subscriptions(ctx context.Context, filters *SubscriptionFilters) ([]Subscription, error)
	CreateSubscription(ctx context.Context, cs *Subscription) error
	UpdateSubscription(ctx context.Context, cs *Subscription) error
	DeleteSubscription(ctx context.Context, cs *Subscription) error
}

type ServiceConfiguration func(svc *Service) error

type Service struct {
	subscriptions SubscriptionRepository
	courses       CourseClient
	matrices      MatrixClient
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

// WithSubscriptionRepository injects the subscription repository to the domain Service.
func WithSubscriptionRepository(sr SubscriptionRepository) ServiceConfiguration {
	return func(svc *Service) error {
		svc.subscriptions = sr

		return nil
	}
}

func WithCourseClient(c CourseClient) ServiceConfiguration {
	return func(svc *Service) error {
		svc.courses = c

		return nil
	}
}

func WithMatrixClient(m MatrixClient) ServiceConfiguration {
	return func(svc *Service) error {
		svc.matrices = m

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
