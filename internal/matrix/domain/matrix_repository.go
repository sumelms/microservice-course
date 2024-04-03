package domain

import "github.com/google/uuid"

type MatrixRepository interface {
	Matrix(matrixUUID uuid.UUID) (Matrix, error)
	CourseMatrix(courseUUID uuid.UUID, matrixUUID uuid.UUID) (Matrix, error)
	Matrices() ([]Matrix, error)
	CourseMatrices(courseUUID uuid.UUID) ([]Matrix, error)
	CreateMatrix(matrix *Matrix) error
	UpdateMatrix(matrix *Matrix) error
	DeleteMatrix(matrix *DeletedMatrix) error
	AddSubject(matrixSubject *MatrixSubject) error
	RemoveSubject(matrixUUID, subjectUUID uuid.UUID) error
}
