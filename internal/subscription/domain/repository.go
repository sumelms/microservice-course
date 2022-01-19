package domain

type Repository interface {
	Subscription(int) (Subscription, error)
	Subscriptions() ([]Subscription, error)
	CreateSubscription(*Subscription) error
	UpdateSubscription(*Subscription) error
	DeleteSubscription(int) error
}
