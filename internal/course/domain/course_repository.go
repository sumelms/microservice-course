package domain

import "github.com/google/uuid"

type CourseRepository interface {
	Course(uuid.UUID) (Course, error)
	Courses() ([]Course, error)
	CreateCourse(*Course) error
	UpdateCourse(*Course) error
	DeleteCourse(uuid.UUID) error
}
