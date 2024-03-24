package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

func NewMatrixRepository(db *sqlx.DB) (MatrixRepository, error) {
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesMatrix() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return MatrixRepository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return MatrixRepository{
		statements: sqlStatements,
	}, nil
}

type MatrixRepository struct {
	statements map[string]*sqlx.Stmt
}

func (r MatrixRepository) statement(s string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[s]
	if !ok {
		return nil, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", s)
	}
	return stmt, nil
}

func (r MatrixRepository) Matrix(matrixUUID uuid.UUID) (domain.Matrix, error) {
	stmt, err := r.statement(getMatrix)
	if err != nil {
		return domain.Matrix{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getMatrix)
	}

	var matrix domain.Matrix
	if err := stmt.Get(&matrix, matrixUUID); err != nil {
		return domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrix")
	}
	return matrix, nil
}

func (r MatrixRepository) CourseMatrix(courseUUID uuid.UUID, matrixUUID uuid.UUID) (domain.Matrix, error) {
	stmt, err := r.statement(getCourseMatrix)
	if err != nil {
		return domain.Matrix{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getMatrix)
	}

	var m domain.Matrix
	if err := stmt.Get(&m, courseUUID, matrixUUID); err != nil {
		return domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrix")
	}
	return m, nil
}

func (r MatrixRepository) Matrices() ([]domain.Matrix, error) {
	stmt, err := r.statement(listMatrix)
	if err != nil {
		return []domain.Matrix{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", listMatrix)
	}

	var list []domain.Matrix
	if err := stmt.Select(&list); err != nil {
		return []domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrices")
	}
	return list, nil
}

func (r MatrixRepository) CourseMatrices(courseUUID uuid.UUID) ([]domain.Matrix, error) {
	stmt, err := r.statement(listCourseMatrices)
	if err != nil {
		return []domain.Matrix{}, errors.NewErrorf(
			errors.ErrCodeUnknown, "prepared statement %s not found", listCourseMatrices)
	}

	var list []domain.Matrix
	if err := stmt.Select(&list, courseUUID); err != nil {
		return []domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrices")
	}
	return list, nil
}

func (r MatrixRepository) CreateMatrix(matrix *domain.Matrix) error {
	stmt, err := r.statement(createMatrix)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createMatrix)
	}

	args := []interface{}{
		matrix.CourseUUID,
		matrix.Code,
		matrix.Name,
		matrix.Description,
	}
	if err := stmt.Get(matrix, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating matrix")
	}
	return nil
}

func (r MatrixRepository) UpdateMatrix(matrix *domain.Matrix) error {
	stmt, err := r.statement(updateMatrix)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateMatrix)
	}

	args := []interface{}{
		matrix.UUID,
		matrix.Code,
		matrix.Name,
		matrix.Description,
	}
	if err := stmt.Get(matrix, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating matrix")
	}
	return nil
}

func (r MatrixRepository) DeleteMatrix(matrixUUID uuid.UUID) error {
	stmt, err := r.statement(deleteMatrix)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteMatrix)
	}

	if _, err := stmt.Exec(matrixUUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting matrix")
	}
	return nil
}

func (r MatrixRepository) AddSubject(ms *domain.MatrixSubject) error {
	stmt, err := r.statement(addSubject)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", addSubject)
	}

	if _, err := stmt.Exec(ms); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error adding subject to matrix")
	}
	return nil
}

func (r MatrixRepository) RemoveSubject(matrixUUID, subjectUUID uuid.UUID) error {
	stmt, err := r.statement(removeSubject)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", removeSubject)
	}

	if _, err := stmt.Exec(matrixUUID, subjectUUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error removing subject from matrix")
	}
	return nil
}
