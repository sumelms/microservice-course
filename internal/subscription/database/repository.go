package database

import (
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type Repository struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewRepository(db *sqlx.DB, l log.Logger) *Repository {
	return &Repository{db: db, logger: l}
}

func (r Repository) Create(subscription *domain.Subscription) (domain.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Find(s string) (domain.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Update(subscription *domain.Subscription) (domain.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Delete(s string) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) List(m map[string]interface{}) ([]domain.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) FindBy(s string, i interface{}) ([]domain.Subscription, error) {
	//TODO implement me
	panic("implement me")
}
