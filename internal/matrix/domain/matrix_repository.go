package domain

import "github.com/google/uuid"

type MatrixRepository interface {
	// CREATE.
	CreateMatrix(matrix *Matrix) error
	CreateMatrixSubject(matrixSubject *MatrixSubject) error
	// READ.
	Matrix(matrixUUID uuid.UUID) (Matrix, error)
	MatrixSubject(matrixSubject *MatrixSubject) error
	CourseMatrixExists(courseUUID uuid.UUID, matrixUUID uuid.UUID) (bool, error)
	Matrices() ([]Matrix, error)
	CourseMatrices(courseUUID uuid.UUID) ([]Matrix, error)
	MatrixSubjects(matrixUUID uuid.UUID) ([]MatrixSubject, error)
	// UPDATE.
	UpdateMatrix(matrix *Matrix) error
	UpdateMatrixSubject(matrixSubject *MatrixSubject) error
	// DELETE.
	DeleteMatrix(matrix *DeletedMatrix) error
	DeleteMatrixSubject(matrix *DeletedMatrixSubject) error
}
