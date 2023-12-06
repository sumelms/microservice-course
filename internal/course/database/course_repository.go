package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

func NewCourseRepository(db *sqlx.DB) (CourseRepository, error) { //nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesCourse() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return CourseRepository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown,
				"error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return CourseRepository{
		statements: sqlStatements,
	}, nil
}

type CourseRepository struct {
	statements map[string]*sqlx.Stmt
}

func (r CourseRepository) statement(s string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[s]
	if !ok {
		return nil, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", s)
	}
	return stmt, nil
}

// Course get the Course by given id.
func (r CourseRepository) Course(id uuid.UUID) (domain.Course, error) {
	stmt, err := r.statement(getCourse)
	if err != nil {
		return domain.Course{}, err
	}

	var c domain.Course
	if err := stmt.Get(&c, id); err != nil {
		return domain.Course{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return c, nil
}

// Courses list all courses.
func (r CourseRepository) Courses() ([]domain.Course, error) {
	stmt, err := r.statement(listCourse)
	if err != nil {
		return []domain.Course{}, err
	}

	var cc []domain.Course
	if err := stmt.Select(&cc); err != nil {
		return []domain.Course{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return cc, nil
}

// CreateCourse creates a new course.
func (r CourseRepository) CreateCourse(c *domain.Course) error {
	stmt, err := r.statement(createCourse)
	if err != nil {
		return err
	}

	args := []interface{}{
		c.Code,
		c.Name,
		c.Underline,
		c.Image,
		c.ImageCover,
		c.Excerpt,
		c.Description,
	}
	if err := stmt.Get(c, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating course")
	}
	return nil
}

// UpdateCourseByID update the given course by ID.
func (r CourseRepository) UpdateCourseByID(c *domain.Course) error {
	stmt, err := r.statement(updateCourseByID)
	if err != nil {
		return err
	}

	args := []interface{}{
		// set
		c.Code,
		c.Name,
		c.Underline,
		c.Image,
		c.ImageCover,
		c.Excerpt,
		c.Description,
		// where
		c.UUID,
	}
	if err := stmt.Get(c, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating course")
	}
	return nil
}

func (r CourseRepository) UpdateCourseByCode(c *domain.Course) error {
	stmt, err := r.statement(updateCourseByCode)
	if err != nil {
		return err
	}

	args := []interface{}{
		// set
		c.Name,
		c.Underline,
		c.Image,
		c.ImageCover,
		c.Excerpt,
		c.Description,
		// where
		c.Code,
	}

	if err := stmt.Get(c, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating course")
	}
	return nil
}

// DeleteCourse soft delete the course by given id.
func (r CourseRepository) DeleteCourse(id uuid.UUID) error {
	stmt, err := r.statement(deleteCourse)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting course")
	}
	return nil
}
