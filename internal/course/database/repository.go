package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

// Repository Course struct
type Repository struct {
	*sqlx.DB
}

// Course get the Course by given id
func (r *Repository) Course(id uuid.UUID) (domain.Course, error) {
	var c domain.Course
	query := `SELECT * FROM courses WHERE deleted_at IS NULL AND uuid = $1`
	if err := r.Get(&c, query, id); err != nil {
		return domain.Course{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return c, nil
}

// Courses list all courses
func (r *Repository) Courses() ([]domain.Course, error) {
	var cc []domain.Course
	query := `SELECT * FROM courses WHERE deleted_at IS NULL`
	if err := r.Select(&cc, query); err != nil {
		return []domain.Course{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting courses")
	}
	return cc, nil
}

// CreateCourse creates a new course
func (r *Repository) CreateCourse(c *domain.Course) error {
	query := `INSERT INTO courses (title, subtitle, excerpt, description) VALUES ($1, $2, $3, $4) RETURNING *`
	if err := r.Get(c, query, c.Title, c.Subtitle, c.Excerpt, c.Description); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating course")
	}
	return nil
}

// UpdateCourse update the given course
func (r *Repository) UpdateCourse(c *domain.Course) error {
	query := `UPDATE courses SET title = $1, subtitle = $2, excerpt = $3, description = $4 WHERE uuid = $5 RETURNING *`
	if err := r.Get(c, query, c.Title, c.Subtitle, c.Excerpt, c.Description, c.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating course")
	}
	return nil
}

// DeleteCourse soft delete the course by given id
func (r *Repository) DeleteCourse(id uuid.UUID) error {
	query := `UPDATE courses SET deleted_at = NOW() WHERE uuid = $1`
	if _, err := r.Exec(query, id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting course")
	}
	return nil
}
