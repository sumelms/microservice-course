package domain

import "github.com/google/uuid"

type Repository interface {
	Create(*Course) error
	Update(*Course) error
	Delete(uuid.UUID) error
	Find(uuid.UUID) (Course, error)
	List() ([]Course, error)
}
