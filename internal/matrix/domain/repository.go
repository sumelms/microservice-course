package domain

import "github.com/google/uuid"

type Repository interface {
	Matrix(id uuid.UUID) (Matrix, error)
	Matrices() ([]Matrix, error)
	CreateMatrix(matrix *Matrix) error
	UpdateMatrix(matrix *Matrix) error
	DeleteMatrix(id uuid.UUID) error
	AddSubject(matrixID, subjectID uuid.UUID) error
	RemoveSubject(matrixID, subjectID uuid.UUID) error
}
