package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subject/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

type Repository struct {
	*sqlx.DB
}

func (r *Repository) Subject(id uuid.UUID) (domain.Subject, error) {
	query := `SELECT * FROM subjects WHERE uuid = $1`

	stmt, err := r.Preparex(query)
	if err != nil {
		return domain.Subject{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement")
	}

	var sub domain.Subject
	if err := stmt.Get(&sub, query, id); err != nil {
		return domain.Subject{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subject")
	}
	return sub, nil
}

func (r *Repository) Subjects() ([]domain.Subject, error) {
	query := `SELECT * FROM subjects`

	stmt, err := r.Preparex(query)
	if err != nil {
		return []domain.Subject{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement")
	}

	var subs []domain.Subject
	if err := stmt.Get(&subs, query); err != nil {
		return []domain.Subject{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subjects")
	}
	return subs, nil
}

func (r *Repository) CreateSubject(sub *domain.Subject) error {
	query := `INSERT INTO subjects (title) VALUES ($1) RETURNING *`

	stmt, err := r.Preparex(query)
	if err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement")
	}

	if err := stmt.Get(sub, query, sub.Title); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating course")
	}
	return nil
}

func (r *Repository) UpdateSubject(sub *domain.Subject) error {
	query := `UPDATE subjects SET title = $1 WHERE uuid = $2 RETURNING *`

	stmt, err := r.Preparex(query)
	if err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement")
	}

	if err := stmt.Get(sub, query, sub.Title, sub.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating course")
	}
	return nil
}

func (r *Repository) DeleteSubject(id uuid.UUID) error {
	query := `UPDATE subjects SET deleted_at = NOW() WHERE uuid = $1`

	stmt, err := r.Preparex(query)
	if err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement")
	}

	if _, err := stmt.Exec(query, id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting course")
	}
	return nil
}
