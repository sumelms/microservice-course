package domain

import (
	"github.com/google/uuid"
)

type Repository interface {
	Subscription(uuid.UUID) (Subscription, error)
	Subscriptions() ([]Subscription, error)
	CreateSubscription(*Subscription) error
	UpdateSubscription(*Subscription) error
	DeleteSubscription(uuid.UUID) error
}
