package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
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

func (r SubscriptionRepository) statement(s string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[s]
	if !ok {
		return nil, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", s)
	}
	return stmt, nil
}

func (r SubscriptionRepository) Subscription(subscriptionUUID uuid.UUID) (domain.Subscription, error) {
	stmt, err := r.statement(getSubscription)
	if err != nil {
		return domain.Subscription{}, err
	}

	var sub domain.Subscription
	if err := stmt.Get(&sub, subscriptionUUID); err != nil {
		return domain.Subscription{},
			errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subscription")
	}
	return sub, nil
}

func (r SubscriptionRepository) Subscriptions() ([]domain.Subscription, error) {
	stmt, err := r.statement(listSubscriptions)
	if err != nil {
		return []domain.Subscription{}, err
	}

	var subs []domain.Subscription
	if err := stmt.Select(&subs); err != nil {
		return []domain.Subscription{},
			errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting subscriptions")
	}
	return subs, nil
}

func (r SubscriptionRepository) CreateSubscriptionWithoutMatrix(sub *domain.Subscription) error {
	stmt, err := r.statement(createSubscriptionWithoutMatrix)
	if err != nil {
		return err
	}

	args := []interface{}{
		sub.CourseUUID,
		sub.UserUUID,
		sub.Role,
		sub.ExpiresAt,
	}
	if err := stmt.Get(sub, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating subscription")
	}
	return nil
}

func (r SubscriptionRepository) CreateSubscription(sub *domain.Subscription) error {
	stmt, err := r.statement(createSubscription)
	if err != nil {
		return err
	}

	args := []interface{}{
		sub.CourseUUID,
		sub.MatrixUUID,
		sub.UserUUID,
		sub.Role,
		sub.ExpiresAt,
	}
	if err := stmt.Get(sub, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating subscription")
	}
	return nil
}

func (r SubscriptionRepository) UpdateSubscription(sub *domain.Subscription) error {
	stmt, err := r.statement(updateSubscription)
	if err != nil {
		return err
	}

	args := []interface{}{
		sub.UUID,
		sub.Role,
		sub.ExpiresAt,
	}
	if _, err := stmt.Exec(args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating subscription")
	}

	updatedSub, err := r.Subscription(sub.UUID)
	*sub = updatedSub

	return err
}

func (r SubscriptionRepository) DeleteSubscription(sub *domain.DeletedSubscription) error {
	stmt, err := r.statement(deleteSubscription)
	if err != nil {
		return err
	}

	args := []interface{}{
		sub.UUID,
		sub.Reason,
	}
	if err := stmt.Get(sub, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting subscription")
	}
	return nil
}

func (r SubscriptionRepository) CourseSubscriptions(courseUUID uuid.UUID) ([]domain.Subscription, error) {
	stmt, err := r.statement(courseSubscriptions)
	if err != nil {
		return []domain.Subscription{}, err
	}

	var subs []domain.Subscription
	if err := stmt.Select(&subs, courseUUID); err != nil {
		return []domain.Subscription{},
			errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course subscriptions")
	}

	return subs, nil
}

func (r SubscriptionRepository) UserSubscriptions(userUUID uuid.UUID) ([]domain.Subscription, error) {
	stmt, err := r.statement(userSubscriptions)
	if err != nil {
		return []domain.Subscription{}, err
	}

	var subs []domain.Subscription
	if err := stmt.Select(&subs, userUUID); err != nil {
		return []domain.Subscription{},
			errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting user subscriptions")
	}

	return subs, nil
}
