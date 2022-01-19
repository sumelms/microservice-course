package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type ServiceInterface interface {
	Subscription(context.Context, int) (Subscription, error)
	Subscriptions(context.Context) ([]Subscription, error)
	CreateSubscription(context.Context, *Subscription) error
	UpdateSubscription(context.Context, *Subscription) error
	DeleteSubscription(context.Context, int) error
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

func (s *Service) Subscriptions(ctx context.Context) ([]Subscription, error) {
	list, err := s.repo.Subscriptions()
	if err != nil {
		return []Subscription{}, fmt.Errorf("service didn't found any subscription: %w", err)
	}
	return list, nil
}

func (s *Service) CreateSubscription(ctx context.Context, sub *Subscription) error {
	if err := s.repo.CreateSubscription(sub); err != nil {
		return fmt.Errorf("service can't create subscription: %w", err)
	}
	return nil
}

func (s *Service) Subscription(ctx context.Context, id int) (Subscription, error) {
	sub, err := s.repo.Subscription(id)
	if err != nil {
		return Subscription{}, fmt.Errorf("service can't find subscription: %w", err)
	}
	return sub, nil
}

func (s *Service) UpdateSubscription(ctx context.Context, sub *Subscription) error {
	if err := s.repo.UpdateSubscription(sub); err != nil {
		return fmt.Errorf("service can't update subscription: %w", err)
	}
	return nil
}

func (s *Service) DeleteSubscription(ctx context.Context, id int) error {
	if err := s.repo.DeleteSubscription(id); err != nil {
		return fmt.Errorf("service can't delete subscription: %w", err)
	}
	return nil
}
