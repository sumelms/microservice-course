package domain

import "github.com/google/uuid"

type Repository interface {
	CreateCourse(*Course) error
	UpdateCourse(*Course) error
	DeleteCourse(uuid.UUID) error
	Course(uuid.UUID) (Course, error)
	Courses() ([]Course, error)
}
