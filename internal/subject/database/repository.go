package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subject/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

// NewRepository creates the subject repository
func NewRepository(db *sqlx.DB) (repository, error) { // nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queries() {
		stmt, err := db.Preparex(string(query))
		if err != nil {
			return repository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return repository{
		statements: sqlStatements,
	}, nil
}

type repository struct {
	statements map[string]*sqlx.Stmt
}

func (r repository) Subject(id uuid.UUID) (domain.Subject, error) {
	stmt, ok := r.statements[getSubject]
	if !ok {
		return domain.Subject{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getSubject)
	}

	var sub domain.Subject
	if err := stmt.Get(&sub, id); err != nil {
		return domain.Subject{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subject")
	}
	return sub, nil
}

func (r repository) Subjects() ([]domain.Subject, error) {
	stmt, ok := r.statements[listSubject]
	if !ok {
		return []domain.Subject{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", listSubject)
	}

	var subs []domain.Subject
	if err := stmt.Select(&subs); err != nil {
		return []domain.Subject{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subjects")
	}
	return subs, nil
}

func (r repository) CreateSubject(sub *domain.Subject) error {
	stmt, ok := r.statements[createSubject]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createSubject)
	}

	if err := stmt.Get(sub, sub.Name); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating subject")
	}
	return nil
}

func (r repository) UpdateSubject(sub *domain.Subject) error {
	stmt, ok := r.statements[updateSubject]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateSubject)
	}

	if err := stmt.Get(sub, sub.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating subject")
	}
	return nil
}

func (r repository) DeleteSubject(id uuid.UUID) error {
	stmt, ok := r.statements[deleteSubject]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteSubject)
	}

	if _, err := stmt.Exec(id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting subject")
	}
	return nil
}
