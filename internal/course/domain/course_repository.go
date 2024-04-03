package domain

import (
	"github.com/google/uuid"
)

type CourseRepository interface {
	Course(courseUUID uuid.UUID) (Course, error)
	Courses() ([]Course, error)
	CreateCourse(course *Course) error
	UpdateCourse(course *Course) error
	DeleteCourse(course *DeletedCourse) error
}
