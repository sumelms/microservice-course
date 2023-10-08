package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

// NewSubscriptionRepository creates the subscription SubscriptionRepository.
func NewSubscriptionRepository(db *sqlx.DB) (SubscriptionRepository, error) { //nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesSubscription() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return SubscriptionRepository{},
				errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return SubscriptionRepository{
		statements: sqlStatements,
	}, nil
}

type SubscriptionRepository struct {
	statements map[string]*sqlx.Stmt
}

func (r SubscriptionRepository) Subscription(id uuid.UUID) (domain.Subscription, error) {
	stmt, ok := r.statements[getSubscription]
	if !ok {
		return domain.Subscription{},
			errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getSubscription)
	}

	var sub domain.Subscription
	if err := stmt.Get(&sub, id); err != nil {
		return domain.Subscription{},
			errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subscription")
	}
	return sub, nil
}

func (r SubscriptionRepository) Subscriptions() ([]domain.Subscription, error) {
	stmt, ok := r.statements[listSubscription]
	if !ok {
		return []domain.Subscription{},
			errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", listSubscription)
	}

	var subs []domain.Subscription
	if err := stmt.Select(&subs); err != nil {
		return []domain.Subscription{},
			errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subscriptions")
	}
	return subs, nil
}

func (r SubscriptionRepository) CreateSubscription(s *domain.Subscription) error {
	stmt, ok := r.statements[createSubscription]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createSubscription)
	}

	args := []interface{}{
		s.CourseID,
		s.MatrixID,
		s.UserID,
		s.ExpiresAt,
	}
	if err := stmt.Get(s, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating subscription")
	}
	return nil
}

func (r SubscriptionRepository) UpdateSubscription(sub *domain.Subscription) error {
	stmt, ok := r.statements[updateSubscription]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateSubscription)
	}

	args := []interface{}{
		sub.UserID,
		sub.CourseID,
		sub.MatrixID,
		sub.ExpiresAt,
		sub.UUID,
	}
	if err := stmt.Get(sub, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating subscription")
	}
	return nil
}

func (r SubscriptionRepository) DeleteSubscription(id uuid.UUID) error {
	stmt, ok := r.statements[deleteSubscription]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteSubscription)
	}

	if _, err := stmt.Exec(id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting subscription")
	}
	return nil
}
