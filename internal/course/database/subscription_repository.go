package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

// NewSubscriptionRepository creates the subscription subscriptionRepository.
func NewSubscriptionRepository(db *sqlx.DB) (subscriptionRepository, error) { //nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesSubscription() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return subscriptionRepository{},
				errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return subscriptionRepository{
		statements: sqlStatements,
	}, nil
}

type subscriptionRepository struct {
	statements map[string]*sqlx.Stmt
}

func (r subscriptionRepository) Subscription(id uuid.UUID) (domain.Subscription, error) {
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

func (r subscriptionRepository) Subscriptions() ([]domain.Subscription, error) {
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

func (r subscriptionRepository) CreateSubscription(s *domain.Subscription) error {
	stmt, ok := r.statements[createSubscription]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createSubscription)
	}

	if err := stmt.Get(s, s.CourseID, s.MatrixID, s.UserID, s.ExpiresAt); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating subscription")
	}
	return nil
}

func (r subscriptionRepository) UpdateSubscription(sub *domain.Subscription) error {
	stmt, ok := r.statements[updateSubscription]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateSubscription)
	}

	if err := stmt.Get(sub, sub.UserID, sub.CourseID, sub.MatrixID, sub.ExpiresAt, sub.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating subscription")
	}
	return nil
}

func (r subscriptionRepository) DeleteSubscription(id uuid.UUID) error {
	stmt, ok := r.statements[deleteSubscription]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteSubscription)
	}

	if _, err := stmt.Exec(id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting subscription")
	}
	return nil
}
