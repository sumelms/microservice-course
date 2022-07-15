package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Subscription(context.Context, uuid.UUID) (Subscription, error)
	Subscriptions(context.Context) ([]Subscription, error)
	CreateSubscription(context.Context, *Subscription) error
	UpdateSubscription(context.Context, *Subscription) error
	DeleteSubscription(context.Context, uuid.UUID) error
}

type CourseService interface {
	ExistCourse(context.Context, uuid.UUID) error
}

type Service struct {
	repo       Repository
	coursesSvc CourseService
	logger     log.Logger
}

func NewService(repo Repository, courseSvc CourseService, logger log.Logger) *Service {
	return &Service{
		repo:       repo,
		logger:     logger,
		coursesSvc: courseSvc,
	}
}

func (s *Service) Subscriptions(_ context.Context) ([]Subscription, error) {
	list, err := s.repo.Subscriptions()
	if err != nil {
		return []Subscription{}, fmt.Errorf("service didn't found any subscription: %w", err)
	}
	return list, nil
}

func (s *Service) CreateSubscription(ctx context.Context, sub *Subscription) error {
	if err := s.coursesSvc.ExistCourse(ctx, sub.CourseID); err != nil {
		return fmt.Errorf("error checking if course %s exists: %w", sub.CourseID, err)
	}
	if err := s.repo.CreateSubscription(sub); err != nil {
		return fmt.Errorf("service can't create subscription: %w", err)
	}
	return nil
}

func (s *Service) Subscription(_ context.Context, id uuid.UUID) (Subscription, error) {
	sub, err := s.repo.Subscription(id)
	if err != nil {
		return Subscription{}, fmt.Errorf("service can't find subscription: %w", err)
	}
	return sub, nil
}

func (s *Service) UpdateSubscription(_ context.Context, sub *Subscription) error {
	if err := s.repo.UpdateSubscription(sub); err != nil {
		return fmt.Errorf("service can't update subscription: %w", err)
	}
	return nil
}

func (s *Service) DeleteSubscription(_ context.Context, id uuid.UUID) error {
	if err := s.repo.DeleteSubscription(id); err != nil {
		return fmt.Errorf("service can't delete subscription: %w", err)
	}
	return nil
}
