package database

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subscription/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

type Repository struct {
	*sqlx.DB
}

func (r *Repository) Subscription(id int) (domain.Subscription, error) {
	var s domain.Subscription
	if err := r.Get(&s, `SELECT * FROM subscriptions WHERE id = $1`, id); err != nil {
		return domain.Subscription{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subscription")
	}
	return s, nil
}

func (r *Repository) Subscriptions() ([]domain.Subscription, error) {
	var ss []domain.Subscription
	if err := r.Select(&ss, `SELECT * FROM subscriptions`); err != nil {
		return []domain.Subscription{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subscriptions")
	}
	return ss, nil
}

func (r *Repository) CreateSubscription(s *domain.Subscription) error {
	if err := r.Get(&s, `INSERT INTO subscriptions VALUES ($1, $2, $3, $4) RETURNING *`,
		s.CourseID,
		s.MatrixID,
		s.UserID,
		s.ValidUntil); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating subscription")
	}
	return nil
}

func (r *Repository) UpdateSubscription(s *domain.Subscription) error {
	if err := r.Get(&s, `UPDATE subscriptions
		SET user_id = $1, course_id = $2, matrix_id = $3, valid_until = $4
		WHERE id = $5
		RETURNING *`,
		s.UserID,
		s.CourseID,
		s.MatrixID,
		s.ValidUntil,
		s.ID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating subscription")
	}
	return nil
}

func (r *Repository) DeleteSubscription(id int) error {
	if _, err := r.Exec(`UPDATE subscriptions SET deleted_at = $1 WHERE id = $2`, time.Now(), id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting subscription")
	}
	return nil
}
