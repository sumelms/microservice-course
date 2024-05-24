package domain

import "github.com/google/uuid"

type MatrixRepository interface {
	// CREATE.
	CreateMatrix(matrix *Matrix) error
	AddSubject(matrixSubject *MatrixSubject) error
	// READ.
	Matrix(matrixUUID uuid.UUID) (Matrix, error)
	CourseMatrixExists(courseUUID uuid.UUID, matrixUUID uuid.UUID) (bool, error)
	Matrices() ([]Matrix, error)
	CourseMatrices(courseUUID uuid.UUID) ([]Matrix, error)
	// UPDATE.
	UpdateMatrix(matrix *Matrix) error
	// DELETE.
	DeleteMatrix(matrix *DeletedMatrix) error
	RemoveSubject(matrixUUID, subjectUUID uuid.UUID) error
}
