package course

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/pkg/errors"
)

type service struct {
	statements map[string]*sqlx.Stmt
}

func NewCourseService(db *sqlx.DB) (service, error) { // nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queries() {
		stmt, err := db.Preparex(string(query))
		if err != nil {
			return service{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return service{
		statements: sqlStatements,
	}, nil
}

func (s service) ExistCourse(_ context.Context, id uuid.UUID) error {
	stmt, ok := s.statements[existCourse]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", existCourse)
	}

	found := stmt.QueryRow(id)
	if found == nil {
		return errors.NewErrorf(errors.ErrCodeNotFound, "course %s not found", id.String())
	}
	return nil
}
