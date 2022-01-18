package database

import (
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/course/domain"
)

type Repository struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewRepository(db *sqlx.DB, l log.Logger) *Repository {
	return &Repository{db: db, logger: l}
}

func (r Repository) Create(course *domain.Course) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Find(s string) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Update(course *domain.Course) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Delete(s string) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) List() ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}
