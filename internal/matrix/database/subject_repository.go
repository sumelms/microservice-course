package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

// NewSubjectRepository creates the subject SubjectRepository.
func NewSubjectRepository(db *sqlx.DB) (SubjectRepository, error) { //nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesSubject() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return SubjectRepository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return SubjectRepository{
		statements: sqlStatements,
	}, nil
}

func (r SubjectRepository) statement(s string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[s]
	if !ok {
		return nil, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", s)
	}
	return stmt, nil
}

type SubjectRepository struct {
	statements map[string]*sqlx.Stmt
}

func (r SubjectRepository) Subject(subjectUUID uuid.UUID) (domain.Subject, error) {
	stmt, err := r.statement(getSubject)
	if err != nil {
		return domain.Subject{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getSubject)
	}

	var subject domain.Subject
	if err := stmt.Get(&subject, subjectUUID); err != nil {
		return domain.Subject{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subject")
	}
	return subject, nil
}

func (r SubjectRepository) Subjects() ([]domain.Subject, error) {
	stmt, err := r.statement(listSubjects)
	if err != nil {
		return []domain.Subject{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", listSubjects)
	}

	var subjects []domain.Subject
	if err := stmt.Select(&subjects); err != nil {
		return []domain.Subject{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subjects")
	}
	return subjects, nil
}

func (r SubjectRepository) CreateSubject(subject *domain.Subject) error {
	stmt, err := r.statement(createSubject)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createSubject)
	}

	args := []interface{}{
		subject.Code,
		subject.Name,
		subject.Objective,
		subject.Credit,
		subject.Workload,
		subject.PublishedAt,
	}
	if err := stmt.Get(subject, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating subject")
	}
	return nil
}

func (r SubjectRepository) UpdateSubject(subject *domain.Subject) error {
	stmt, err := r.statement(updateSubject)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateSubject)
	}

	args := []interface{}{
		subject.UUID,
		subject.Code,
		subject.Name,
		subject.Objective,
		subject.Credit,
		subject.Workload,
		subject.PublishedAt,
	}
	if err := stmt.Get(subject, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating subject")
	}
	return nil
}

func (r SubjectRepository) DeleteSubject(subject *domain.DeletedSubject) error {
	stmt, err := r.statement(deleteSubject)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteSubject)
	}

	args := []interface{}{
		subject.UUID,
	}
	if err := stmt.Get(subject, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting subject")
	}
	return nil
}
