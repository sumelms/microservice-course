package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Subscription(_ context.Context, subscriptionUUID uuid.UUID) (Subscription, error) {
	sub, err := s.subscriptions.Subscription(subscriptionUUID)
	if err != nil {
		return Subscription{}, fmt.Errorf("service can't find subscription: %w", err)
	}

	return sub, nil
}

func (s *Service) Subscriptions(_ context.Context, filters *SubscriptionFilters) ([]Subscription, error) {
	list, err := func() ([]Subscription, error) {
		if filters != nil {
			if filters.UserUUID != uuid.Nil {
				return s.subscriptions.UserSubscriptions(filters.UserUUID)
			}

			if filters.CourseUUID != uuid.Nil {
				return s.subscriptions.CourseSubscriptions(filters.CourseUUID)
			}
		}

		return s.subscriptions.Subscriptions()
	}()
	if err != nil {
		return []Subscription{}, fmt.Errorf("service didn't found any subscription: %w", err)
	}

	return list, nil
}

func (s *Service) CreateSubscription(ctx context.Context, sub *Subscription) error {
	err := s.courses.CourseExists(ctx, *sub.CourseUUID)
	if err != nil {
		return fmt.Errorf("error checking if course %s exists: %w", sub.CourseUUID, err)
	}

	if sub.MatrixUUID == nil {
		if err := s.subscriptions.CreateSubscriptionWithoutMatrix(sub); err != nil {
			return fmt.Errorf("service can't create subscription: %w", err)
		}

		return nil
	}

	if err := s.matrices.CourseMatrixExists(ctx, *sub.CourseUUID, *sub.MatrixUUID); err != nil {
		return fmt.Errorf("error checking if matrix %s exists: %w", sub.MatrixUUID, err)
	}

	if err := s.subscriptions.CreateSubscription(sub); err != nil {
		return fmt.Errorf("service can't create subscription: %w", err)
	}

	return nil
}

func (s *Service) UpdateSubscription(_ context.Context, subscription *Subscription) (Subscription, error) {
	sub, err := s.subscriptions.UpdateSubscription(subscription)
	if err != nil {
		return Subscription{}, fmt.Errorf("service can't update subscription: %w", err)
	}

	return sub, nil
}

func (s *Service) DeleteSubscription(_ context.Context, sub *DeletedSubscription) error {
	if err := s.subscriptions.DeleteSubscription(sub); err != nil {
		return fmt.Errorf("service can't delete subscription: %w", err)
	}

	return nil
}
