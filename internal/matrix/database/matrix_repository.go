package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

// NewMatrixRepository creates the matrix matrixRepository
func NewMatrixRepository(db *sqlx.DB) (matrixRepository, error) {
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesMatrix() {
		stmt, err := db.Preparex(string(query))
		if err != nil {
			return matrixRepository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return matrixRepository{
		statements: sqlStatements,
	}, nil
}

type matrixRepository struct {
	statements map[string]*sqlx.Stmt
}

// Matrix get the matrix by given id
func (r matrixRepository) Matrix(id uuid.UUID) (domain.Matrix, error) {
	stmt, ok := r.statements[getMatrix]
	if !ok {
		return domain.Matrix{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getMatrix)
	}

	var m domain.Matrix
	if err := stmt.Get(&m, id); err != nil {
		return domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrix")
	}
	return m, nil
}

// Matrices get the list of matrices
func (r matrixRepository) Matrices() ([]domain.Matrix, error) {
	stmt, ok := r.statements[listMatrix]
	if !ok {
		return []domain.Matrix{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", listMatrix)
	}

	var mm []domain.Matrix
	if err := stmt.Select(&mm); err != nil {
		return []domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrices")
	}
	return mm, nil
}

// CreateMatrix create a new matrix
func (r matrixRepository) CreateMatrix(m *domain.Matrix) error {
	stmt, ok := r.statements[createMatrix]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createMatrix)
	}

	if err := stmt.Get(m, m.Code, m.Name, m.Description, m.CourseID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating matrix")
	}
	return nil
}

// UpdateMatrix updates the given matrix
func (r matrixRepository) UpdateMatrix(m *domain.Matrix) error {
	stmt, ok := r.statements[updateMatrix]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateMatrix)
	}

	if err := stmt.Get(m, m.Code, m.Name, m.Description, m.CourseID, m.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating matrix")
	}
	return nil
}

// DeleteMatrix delete the given matrix by uuid
func (r matrixRepository) DeleteMatrix(id uuid.UUID) error {
	stmt, ok := r.statements[deleteMatrix]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteMatrix)
	}

	if _, err := stmt.Exec(id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting matrix")
	}
	return nil
}

// AddSubject adds the subject to the matrix
func (r matrixRepository) AddSubject(matrixID, subjectID uuid.UUID) error {
	stmt, ok := r.statements[addSubject]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", addSubject)
	}

	if _, err := stmt.Exec(matrixID, subjectID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error adding subject to matrix")
	}
	return nil
}

// RemoveSubject removes the subject from the matrix
func (r matrixRepository) RemoveSubject(matrixID, subjectID uuid.UUID) error {
	stmt, ok := r.statements[removeSubject]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", removeSubject)
	}

	if _, err := stmt.Exec(matrixID, subjectID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error removing subject from matrix")
	}
	return nil
}
