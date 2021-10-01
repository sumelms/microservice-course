package domain

type Repository interface {
	Create(*Subscription) (Subscription, error)
	Find(string) (Subscription, error)
	Update(*Subscription) (Subscription, error)
	Delete(string) error
	List() ([]Subscription, error)
}
