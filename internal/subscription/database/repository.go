package database

import (
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
	merrors "github.com/sumelms/microservice-course/pkg/errors"
)

const (
	whereSubscriptionID = "ID = ?"
)

type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&Subscription{})

	return &Repository{db: db, logger: logger}
}

func (r *Repository) Create(subscription *domain.Subscription) (domain.Subscription, error) {
	entity := toDBModel(subscription)

	if err := r.db.Create(&entity).Error; err != nil {
		return domain.Subscription{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't create subscription")
	}
	return toDomainModel(&entity), nil
}

func (r *Repository) Find(id string) (domain.Subscription, error) {
	var subscription Subscription

	query := r.db.Where(whereSubscriptionID, id).First(&subscription)
	if query.RecordNotFound() {
		return domain.Subscription{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "subscription not found")
	}
	if err := query.Error; err != nil {
		return domain.Subscription{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "find subscription")
	}

	return toDomainModel(&subscription), nil
}

func (r *Repository) Update(s *domain.Subscription) (domain.Subscription, error) {
	var subscription Subscription

	query := r.db.Where(whereSubscriptionID, s.ID).First(&subscription)

	if query.RecordNotFound() {
		return domain.Subscription{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "subscription not found")
	}

	query = r.db.Model(&subscription).Updates(&s)

	if err := query.Error; err != nil {
		return domain.Subscription{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't update subscription")
	}

	return *s, nil
}

func (r *Repository) Delete(id string) error {
	query := r.db.Where(whereSubscriptionID, id).Delete(&Subscription{})

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merrors.WrapErrorf(err, merrors.ErrCodeNotFound, "subscription not found")
		}
	}

	return nil
}

func (r *Repository) List() ([]domain.Subscription, error) {
	var subscriptions []Subscription

	query := r.db.Find(&subscriptions)
	if query.RecordNotFound() {
		return []domain.Subscription{}, nil
	}

	var list []domain.Subscription
	for i := range subscriptions {
		s := subscriptions[i]
		list = append(list, toDomainModel(&s))
	}

	return list, nil
}

func (r *Repository) FindBy(field string, value interface{}) ([]domain.Subscription, error) {
	var subscriptions []Subscription

	where := fmt.Sprintf("%s = ?", field)
	query := r.db.Where(where, value).Find(&subscriptions)
	if query.RecordNotFound() {
		return []domain.Subscription{}, nil
	}
	if err := query.Error; err != nil {
		return []domain.Subscription{}, merrors.WrapErrorf(err, merrors.ErrCodeNotFound, fmt.Sprintf("find subscriptions by %s", field))
	}

	var list []domain.Subscription
	for i := range subscriptions {
		s := subscriptions[i]
		list = append(list, toDomainModel(&s))
	}

	return list, nil
}
