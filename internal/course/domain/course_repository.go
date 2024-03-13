package domain

import "github.com/google/uuid"

type CourseRepository interface {
	Course(courseUUID uuid.UUID) (Course, error)
	Courses() ([]Course, error)
	CreateCourse(c *Course) error
	UpdateCourse(c *Course) error
	DeleteCourse(courseUUID uuid.UUID) error
}
