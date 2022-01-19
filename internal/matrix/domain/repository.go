package domain

import "github.com/google/uuid"

type Repository interface {
	Matrix(uuid.UUID) (Matrix, error)
	Matrices() ([]Matrix, error)
	CreateMatrix(*Matrix) error
	UpdateMatrix(*Matrix) error
	DeleteMatrix(uuid.UUID) error
	FindMatricesBy(string, interface{}) ([]Matrix, error)
}
