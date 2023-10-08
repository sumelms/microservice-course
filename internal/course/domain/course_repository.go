package domain

import "github.com/google/uuid"

type CourseRepository interface {
	Course(id uuid.UUID) (Course, error)
	Courses() ([]Course, error)
	CreateCourse(c *Course) error
	UpdateCourseByID(c *Course) error
	UpdateCourseByCode(c *Course) error
	DeleteCourse(id uuid.UUID) error
}
