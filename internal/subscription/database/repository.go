package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subscription/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

type Repository struct {
	*sqlx.DB
}

func (r *Repository) Subscription(id uuid.UUID) (domain.Subscription, error) {
	var s domain.Subscription
	query := `SELECT * FROM subscriptions WHERE deleted_at IS NULL AND uuid = $1`
	if err := r.Get(&s, query, id); err != nil {
		return domain.Subscription{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subscription")
	}
	return s, nil
}

func (r *Repository) Subscriptions() ([]domain.Subscription, error) {
	var ss []domain.Subscription
	query := `SELECT * FROM subscriptions WHERE deleted_at IS NULL`
	if err := r.Select(&ss, query); err != nil {
		return []domain.Subscription{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subscriptions")
	}
	return ss, nil
}

func (r *Repository) CreateSubscription(s *domain.Subscription) error {
	query := `INSERT INTO subscriptions (course_id, matrix_id, user_id, valid_until) VALUES ($1, $2, $3, $4) RETURNING *`
	if err := r.Get(&s, query, s.CourseID, s.MatrixID, s.UserID, s.ValidUntil); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating subscription")
	}
	return nil
}

func (r *Repository) UpdateSubscription(s *domain.Subscription) error {
	query := `UPDATE subscriptions SET user_id = $1, course_id = $2, matrix_id = $3, valid_until = $4 WHERE uuid = $5 RETURNING *`
	if err := r.Get(&s, query, s.UserID, s.CourseID, s.MatrixID, s.ValidUntil, s.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating subscription")
	}
	return nil
}

func (r *Repository) DeleteSubscription(id uuid.UUID) error {
	query := `UPDATE subscriptions SET deleted_at = $1 WHERE uuid = $2`
	if _, err := r.Exec(query, time.Now(), id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting subscription")
	}
	return nil
}
