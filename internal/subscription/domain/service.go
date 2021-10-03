package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type ServiceInterface interface {
	ListSubscription(context.Context) ([]Subscription, error)
	CreateSubscription(context.Context, *Subscription) (Subscription, error)
	FindSubscription(context.Context, string) (Subscription, error)
	UpdateSubscription(context.Context, *Subscription) (Subscription, error)
	DeleteSubscription(context.Context, string) error
	FindSubscriptionByCourse(context.Context, string) ([]Subscription, error)
	FindSubscriptionByUser(context.Context, string) ([]Subscription, error)
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

func (s *Service) ListSubscription(ctx context.Context) ([]Subscription, error) {
	list, err := s.repo.List()
	if err != nil {
		return []Subscription{}, fmt.Errorf("Service didn't found any subscription: %w", err)
	}
	return list, nil
}

func (s *Service) CreateSubscription(ctx context.Context, subscription *Subscription) (Subscription, error) {
	sub, err := s.repo.Create(subscription)
	if err != nil {
		return Subscription{}, fmt.Errorf("Service can't create subscription: %w", err)
	}
	return sub, nil
}

func (s *Service) FindSubscription(ctx context.Context, id string) (Subscription, error) {
	sub, err := s.repo.Find(id)
	if err != nil {
		return Subscription{}, fmt.Errorf("Service can't find subscription: %w", err)
	}
	return sub, nil
}

func (s *Service) UpdateSubscription(ctx context.Context, subscription *Subscription) (Subscription, error) {
	sub, err := s.repo.Update(subscription)
	if err != nil {
		return Subscription{}, fmt.Errorf("Service can't update subscription: %w", err)
	}
	return sub, nil
}

func (s *Service) DeleteSubscription(ctx context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("Service can't delete subscription: %w", err)
	}
	return nil
}

func (s *Service) FindSubscriptionByCourse(ctx context.Context, id string) ([]Subscription, error) {
	list, err := s.repo.FindBy("course_id", id)
	if err != nil {
		return []Subscription{}, fmt.Errorf("Service can't find subscriptions to course %s: %w", id, err)
	}
	return list, nil
}

func (s *Service) FindSubscriptionByUser(ctx context.Context, id string) ([]Subscription, error) {
	list, err := s.repo.FindBy("user_id", id)
	if err != nil {
		return []Subscription{}, fmt.Errorf("Service can't find subscriptions to user %s: %w", id, err)
	}
	return list, nil
}
