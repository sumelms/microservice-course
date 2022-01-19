package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

type Repository struct {
	*sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{DB: db}
}

func (r Repository) CreateMatrix(m *domain.Matrix) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) UpdateMatrix(m *domain.Matrix) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) DeleteMatrix(id uuid.UUID) error {
	if _, err := r.Exec(`UPDATE matrices SET deleted_at = $1 WHERE id = $2`, time.Now(), id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting course")
	}
	return nil
}

func (r Repository) Matrix(id uuid.UUID) (domain.Matrix, error) {
	var m domain.Matrix
	if err := r.Get(&m, `SELECT * FROM matrices WHERE uuid = $1`, id); err != nil {
		return domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return m, nil
}

func (r Repository) Matrices() ([]domain.Matrix, error) {
	var mm []domain.Matrix
	if err := r.Select(&mm, `SELECT * FROM matrices`); err != nil {
		return []domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrices")
	}
	return mm, nil
}

func (r Repository) FindMatricesBy(s string, i interface{}) ([]domain.Matrix, error) {
	//TODO implement me
	panic("implement me")
}
