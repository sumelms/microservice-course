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

func (r MatrixRepository) CourseMatrixExists(courseUUID uuid.UUID, matrixUUID uuid.UUID) (bool, error) {
	stmt, err := r.statement(getCourseMatrixExists)
	if err != nil {
		return false, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getMatrix)
	}

	var exists bool
	if err := stmt.Get(&exists, courseUUID, matrixUUID); err != nil {
		return false, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrix")
	}
	return exists, nil
}

func (r MatrixRepository) Matrices() ([]domain.Matrix, error) {
	stmt, err := r.statement(listMatrices)
	if err != nil {
		return []domain.Matrix{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", listMatrices)
	}

	var matrices []domain.Matrix
	if err := stmt.Select(&matrices); err != nil {
		return []domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrices")
	}
	return matrices, nil
}

func (r MatrixRepository) CourseMatrices(courseUUID uuid.UUID) ([]domain.Matrix, error) {
	stmt, err := r.statement(listCourseMatrices)
	if err != nil {
		return []domain.Matrix{}, errors.NewErrorf(
			errors.ErrCodeUnknown, "prepared statement %s not found", listCourseMatrices)
	}

	var matrices []domain.Matrix
	if err := stmt.Select(&matrices, courseUUID); err != nil {
		return []domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrices")
	}
	return matrices, nil
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
	if _, err := stmt.Exec(args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating matrix")
	}

	updatedMatrix, err := r.Matrix(matrix.UUID)
	*matrix = updatedMatrix

	return err
}

func (r MatrixRepository) DeleteMatrix(matrix *domain.DeletedMatrix) error {
	stmt, err := r.statement(deleteMatrix)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteMatrix)
	}

	args := []interface{}{
		matrix.UUID,
	}
	if err := stmt.Get(matrix, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting matrix")
	}
	return nil
}

func (r MatrixRepository) CreateMatrixSubject(matrixSubject *domain.MatrixSubject) error {
	stmt, err := r.statement(createMatrixSubject)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createMatrixSubject)
	}

	args := []interface{}{
		matrixSubject.MatrixUUID,
		matrixSubject.SubjectUUID,
		matrixSubject.IsRequired,
		matrixSubject.Group,
	}

	if err := stmt.Get(matrixSubject, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating matrix subject")
	}

	return nil
}

func (r MatrixRepository) MatrixSubjects(matrixUUID uuid.UUID) ([]domain.MatrixSubject, error) {
	stmt, err := r.statement(listMatrixSubjects)
	if err != nil {
		return []domain.MatrixSubject{}, errors.NewErrorf(
			errors.ErrCodeUnknown, "prepared statement %s not found", listMatrixSubjects)
	}

	var matrixSubjects []domain.MatrixSubject
	if err := stmt.Select(&matrixSubjects, matrixUUID); err != nil {
		return []domain.MatrixSubject{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrix subjects")
	}
	return matrixSubjects, nil
}

func (r MatrixRepository) MatrixSubject(matrixSubject *domain.MatrixSubject) error {
	stmt, err := r.statement(getMatrixSubject)
	if err != nil {
		return errors.NewErrorf(
			errors.ErrCodeUnknown, "prepared statement %s not found", getMatrixSubject)
	}

	args := []interface{}{
		matrixSubject.MatrixUUID,
		matrixSubject.SubjectUUID,
	}
	if err := stmt.Get(matrixSubject, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrix subject")
	}
	return nil
}

func (r MatrixRepository) UpdateMatrixSubject(matrixSubject *domain.MatrixSubject) error {
	stmt, err := r.statement(updateMatrixSubject)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateMatrixSubject)
	}

	args := []interface{}{
		matrixSubject.MatrixUUID,
		matrixSubject.SubjectUUID,
		matrixSubject.IsRequired,
		matrixSubject.Group,
	}
	if err := stmt.Get(matrixSubject, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating matrix subject")
	}

	return nil
}

func (r MatrixRepository) DeleteMatrixSubject(matrixSubject *domain.DeletedMatrixSubject) error {
	stmt, err := r.statement(deleteMatrixSubject)
	if err != nil {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteMatrixSubject)
	}

	args := []interface{}{
		matrixSubject.MatrixUUID,
		matrixSubject.SubjectUUID,
	}
	if err := stmt.Get(matrixSubject, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error removing subject from matrix")
	}
	return nil
}
