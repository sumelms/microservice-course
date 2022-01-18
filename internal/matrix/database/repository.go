package database

import (
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type Repository struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewRepository(db *sqlx.DB, l log.Logger) *Repository {
	return &Repository{db: db, logger: l}
}

func (r Repository) Create(matrix *domain.Matrix) (domain.Matrix, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Find(s string) (domain.Matrix, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Update(matrix *domain.Matrix) (domain.Matrix, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Delete(s string) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) List(m map[string]interface{}) ([]domain.Matrix, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) FindBy(s string, i interface{}) ([]domain.Matrix, error) {
	//TODO implement me
	panic("implement me")
}
