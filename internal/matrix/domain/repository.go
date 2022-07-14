package domain

import "github.com/google/uuid"

type Repository interface {
	Matrix(id uuid.UUID) (Matrix, error)
	Matrices() ([]Matrix, error)
	CreateMatrix(matrix *Matrix) error
	UpdateMatrix(matrix *Matrix) error
	DeleteMatrix(id uuid.UUID) error
	// @TODO Can I make these two methods as an separated interface?
	AddSubject(matrixID, subjectID uuid.UUID) error
	RemoveSubject(matrixID, subjectID uuid.UUID) error
}
