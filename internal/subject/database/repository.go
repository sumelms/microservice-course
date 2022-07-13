package database

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subject/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

const (
	listSubject    = "list subject"
	getSubject     = "get subject by uuid"
	delecteSubject = "delete subject by uuid"
)

func NewRepository(db *sqlx.DB) (repository, error) {
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queries() {
		stmt, err := db.Preparex(string(query))
		if err != nil {
			return repository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}
	fmt.Println(sqlStatements)
	return repository{
		statements: sqlStatements,
	}, nil
}

type repository struct {
	statements map[string]*sqlx.Stmt
}

type Query string

func queries() map[string]Query {
	return map[string]Query{
		listSubject:    Query("SELECT * FROM subjects"),
		getSubject:     Query("SELECT * FROM subjects WHERE uuid = $1"),
		delecteSubject: Query("UPDATE subjects SET deleted_at = NOW() WHERE uuid = $1"),
	}
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
	/*
		query := `INSERT INTO subjects (title) VALUES ($1) RETURNING *`

		stmt, err := r.Preparex(query)
		if err != nil {
			return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement")
		}

		if err := stmt.Get(sub, query, sub.Title); err != nil {
			return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating course")
		}
	*/
	return nil
}

func (r repository) UpdateSubject(sub *domain.Subject) error {
	/*
		query := `UPDATE subjects SET title = $1 WHERE uuid = $2 RETURNING *`

		stmt, err := r.Preparex(query)
		if err != nil {
			return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement")
		}

		if err := stmt.Get(sub, query, sub.Title, sub.UUID); err != nil {
			return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating course")
		}
	*/
	return nil
}

func (r repository) DeleteSubject(id uuid.UUID) error {
	stmt, ok := r.statements[delecteSubject]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", delecteSubject)
	}

	if _, err := stmt.Exec(id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting course")
	}
	return nil
}
