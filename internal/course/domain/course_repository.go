package domain

import "github.com/google/uuid"

type CourseRepository interface {
	Course(id uuid.UUID) (Course, error)
	Courses() ([]Course, error)
	CreateCourse(lesson *Course) error
	UpdateCourse(lesson *Course) error
	DeleteCourse(id uuid.UUID) error
}
