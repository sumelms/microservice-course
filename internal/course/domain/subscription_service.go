package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Subscription(_ context.Context, id uuid.UUID) (Subscription, error) {
	sub, err := s.subscriptions.Subscription(id)
	if err != nil {
		return Subscription{}, fmt.Errorf("Service can't find subscription: %w", err)
	}
	return sub, nil
}

func (s *Service) Subscriptions(_ context.Context) ([]Subscription, error) {
	list, err := s.subscriptions.Subscriptions()
	if err != nil {
		return []Subscription{}, fmt.Errorf("Service didn't found any subscription: %w", err)
	}
	return list, nil
}

func (s *Service) CreateSubscription(_ context.Context, sub *Subscription) error {
	_, err := s.courses.Course(sub.CourseID)
	if err != nil {
		return fmt.Errorf("error checking if course %s exists: %w", sub.CourseID, err)
	}
	if err := s.subscriptions.CreateSubscription(sub); err != nil {
		return fmt.Errorf("Service can't create subscription: %w", err)
	}
	return nil
}

func (s *Service) UpdateSubscription(_ context.Context, sub *Subscription) error {
	if err := s.subscriptions.UpdateSubscription(sub); err != nil {
		return fmt.Errorf("Service can't update subscription: %w", err)
	}
	return nil
}

func (s *Service) DeleteSubscription(_ context.Context, id uuid.UUID) error {
	if err := s.subscriptions.DeleteSubscription(id); err != nil {
		return fmt.Errorf("Service can't delete subscription: %w", err)
	}
	return nil
}
