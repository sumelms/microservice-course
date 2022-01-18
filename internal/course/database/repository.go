package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

type Repository struct {
	*sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(c *domain.Course) error {
	if err := r.Get(&c, `INSERT INTO courses VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		c.Title,
		c.Subtitle,
		c.Excerpt,
		c.Description); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating course")
	}
	return nil
}

func (r *Repository) Update(c *domain.Course) error {
	if err := r.Get(&c, `UPDATE courses 
		SET title = $1, subtitle = $2, excerpt = $3, description = $4 
		WHERE uuid = $5 
		RETURNING *`,
		c.Title,
		c.Subtitle,
		c.Excerpt,
		c.Description,
		c.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating course")
	}
	return nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	if _, err := r.Query(`UPDATE courses SET deleted_at = $1 WHERE id = $2`, time.Now(), id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting course")
	}
	return nil
}

func (r *Repository) Find(id uuid.UUID) (domain.Course, error) {
	var c domain.Course
	if err := r.Get(&c, `SELECT * FROM courses WHERE uuid = $1`, id); err != nil {
		return domain.Course{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return c, nil
}

func (r *Repository) List() ([]domain.Course, error) {
	var cc []domain.Course
	if err := r.Get(&cc, `SELECT * FROM courses`); err != nil {
		return []domain.Course{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting courses")
	}
	return cc, nil
}
