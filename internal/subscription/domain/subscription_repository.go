package domain

import "github.com/google/uuid"

type SubscriptionRepository interface {
	Subscription(subscriptionUUID uuid.UUID) (Subscription, error)
	Subscriptions() ([]Subscription, error)
	CreateSubscription(subscription *Subscription) error
	CreateSubscriptionWithoutMatrix(subscription *Subscription) error
	UpdateSubscription(subscription *Subscription) (Subscription, error)
	DeleteSubscription(subscription *DeletedSubscription) error
	CourseSubscriptions(courseUUID uuid.UUID) ([]Subscription, error)
	UserSubscriptions(userUUID uuid.UUID) ([]Subscription, error)
}
