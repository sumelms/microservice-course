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

func (r *Repository) Matrix(id uuid.UUID) (domain.Matrix, error) {
	var m domain.Matrix
	if err := r.Get(&m, `SELECT * FROM matrices WHERE uuid = $1`, id); err != nil {
		return domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return m, nil
}

func (r *Repository) Matrices() ([]domain.Matrix, error) {
	var mm []domain.Matrix
	if err := r.Select(&mm, `SELECT * FROM matrices`); err != nil {
		return []domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrices")
	}
	return mm, nil
}

func (r *Repository) CreateMatrix(m *domain.Matrix) error {
	if err := r.Get(&m, `INSERT INTO subscriptions VALUES ($1, $2, $3) RETURNING *`,
		m.Title,
		m.Description,
		m.CourseID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating matrix")
	}
	return nil
}

func (r *Repository) UpdateMatrix(m *domain.Matrix) error {
	if err := r.Get(&m, `UPDATE matrices
		SET title = $1, description = $2, course_id = $3
		WHERE id = $4`); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating subscription")
	}
	return nil
}

func (r *Repository) DeleteMatrix(id uuid.UUID) error {
	if _, err := r.Exec(`UPDATE matrices SET deleted_at = $1 WHERE id = $2`, time.Now(), id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting course")
	}
	return nil
}
