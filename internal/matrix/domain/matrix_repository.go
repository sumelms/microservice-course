package domain

import "github.com/google/uuid"

type MatrixRepository interface {
	Matrix(id uuid.UUID) (Matrix, error)
	CourseMatrix(courseUUID uuid.UUID, matrixUUID uuid.UUID) (Matrix, error)
	Matrices() ([]Matrix, error)
	CreateMatrix(matrix *Matrix) error
	UpdateMatrix(matrix *Matrix) error
	DeleteMatrix(id uuid.UUID) error
	AddSubject(matrixSubject *MatrixSubject) error
	RemoveSubject(matrixID, subjectID uuid.UUID) error
}
