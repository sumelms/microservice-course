package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

func NewCourseRepository(db *sqlx.DB) (courseRepository, error) { // nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range courseQueries() {
		stmt, err := db.Preparex(string(query))
		if err != nil {
			return courseRepository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return courseRepository{
		statements: sqlStatements,
	}, nil
}

type courseRepository struct {
	statements map[string]*sqlx.Stmt
}

// Course get the Course by given id
func (r courseRepository) Course(id uuid.UUID) (domain.Course, error) {
	stmt, ok := r.statements[getCourse]
	if !ok {
		return domain.Course{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getCourse)
	}

	var c domain.Course
	if err := stmt.Get(&c, id); err != nil {
		return domain.Course{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return c, nil
}

// Courses list all courses
func (r courseRepository) Courses() ([]domain.Course, error) {
	stmt, ok := r.statements[listCourse]
	if !ok {
		return []domain.Course{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", listCourse)
	}

	var cc []domain.Course
	if err := stmt.Select(&cc); err != nil {
		return []domain.Course{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return cc, nil
}

// CreateCourse creates a new course
func (r courseRepository) CreateCourse(c *domain.Course) error {
	stmt, ok := r.statements[createCourse]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createCourse)
	}

	if err := stmt.Get(c, c.Name, c.Underline, c.Image, c.ImageCover, c.Excerpt, c.Description); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating course")
	}
	return nil
}

// UpdateCourse update the given course
func (r courseRepository) UpdateCourse(c *domain.Course) error {
	stmt, ok := r.statements[updateCourse]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateCourse)
	}

	if err := stmt.Get(c, c.Name, c.Underline, c.Image, c.ImageCover, c.Excerpt, c.Description, c.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating course")
	}
	return nil
}

// DeleteCourse soft delete the course by given id
func (r courseRepository) DeleteCourse(id uuid.UUID) error {
	stmt, ok := r.statements[deleteCourse]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteCourse)
	}

	if _, err := stmt.Exec(id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting course")
	}
	return nil
}
