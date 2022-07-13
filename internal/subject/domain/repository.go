package domain

import "github.com/google/uuid"

type Repository interface {
	Subject(uuid.UUID) (Subject, error)
	Subjects() ([]Subject, error)
	CreateSubject(*Subject) error
	UpdateSubject(*Subject) error
	DeleteSubject(uuid.UUID) error
}
