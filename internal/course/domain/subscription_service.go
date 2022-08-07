package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *service) Subscriptions(_ context.Context) ([]Subscription, error) {
	list, err := s.subscriptions.Subscriptions()
	if err != nil {
		return []Subscription{}, fmt.Errorf("service didn't found any subscription: %w", err)
	}
	return list, nil
}

func (s *service) CreateSubscription(_ context.Context, sub *Subscription) error {
	_, err := s.courses.Course(sub.CourseID)
	if err != nil {
		return fmt.Errorf("error checking if course %s exists: %w", sub.CourseID, err)
	}
	if err := s.subscriptions.CreateSubscription(sub); err != nil {
		return fmt.Errorf("service can't create subscription: %w", err)
	}
	return nil
}

func (s *service) Subscription(_ context.Context, id uuid.UUID) (Subscription, error) {
	sub, err := s.subscriptions.Subscription(id)
	if err != nil {
		return Subscription{}, fmt.Errorf("service can't find subscription: %w", err)
	}
	return sub, nil
}

func (s *service) UpdateSubscription(_ context.Context, sub *Subscription) error {
	if err := s.subscriptions.UpdateSubscription(sub); err != nil {
		return fmt.Errorf("service can't update subscription: %w", err)
	}
	return nil
}

func (s *service) DeleteSubscription(_ context.Context, id uuid.UUID) error {
	if err := s.subscriptions.DeleteSubscription(id); err != nil {
		return fmt.Errorf("service can't delete subscription: %w", err)
	}
	return nil
}
