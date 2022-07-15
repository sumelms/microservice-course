package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subscription/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

// NewRepository creates the subscription repository
func NewRepository(db *sqlx.DB) (repository, error) { // nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queries() {
		stmt, err := db.Preparex(string(query))
		if err != nil {
			return repository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return repository{
		statements: sqlStatements,
	}, nil
}

type repository struct {
	statements map[string]*sqlx.Stmt
}

func (r repository) Subscription(id uuid.UUID) (domain.Subscription, error) {
	stmt, ok := r.statements[getSubscription]
	if !ok {
		return domain.Subscription{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", getSubscription)
	}

	var sub domain.Subscription
	if err := stmt.Get(&sub, id); err != nil {
		return domain.Subscription{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subscription")
	}
	return sub, nil
}

func (r repository) Subscriptions() ([]domain.Subscription, error) {
	stmt, ok := r.statements[listSubscription]
	if !ok {
		return []domain.Subscription{}, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", listSubscription)
	}

	var subs []domain.Subscription
	if err := stmt.Select(&subs); err != nil {
		return []domain.Subscription{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subscriptions")
	}
	return subs, nil
}

func (r repository) CreateSubscription(s *domain.Subscription) error {
	stmt, ok := r.statements[createSubscription]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", createSubscription)
	}

	if err := stmt.Get(s, s.CourseID, s.MatrixID, s.UserID, s.ValidUntil); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating subscription")
	}
	return nil
}

func (r repository) UpdateSubscription(sub *domain.Subscription) error {
	stmt, ok := r.statements[updateSubscription]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", updateSubscription)
	}

	if err := stmt.Get(sub, sub.UserID, sub.CourseID, sub.MatrixID, sub.ValidUntil, sub.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating subscription")
	}
	return nil
}

func (r repository) DeleteSubscription(id uuid.UUID) error {
	stmt, ok := r.statements[deleteSubscription]
	if !ok {
		return errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", deleteSubscription)
	}

	if _, err := stmt.Exec(id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting subscription")
	}
	return nil
}
