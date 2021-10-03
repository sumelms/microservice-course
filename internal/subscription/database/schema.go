package database

import (
	"time"

	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type Subscription struct {
	gorm.Model
	UserID     uuid.UUID  `gorm:"type:uuid" sql:"index"`
	CourseID   uuid.UUID  `gorm:"type:uuid" sql:"index"`
	MatrixID   uuid.UUID  `gorm:"type:uuid" sql:"index"`
	ValidUntil *time.Time `sql:"index"`
}

func (s *Subscription) BeforeCreate(scope *gorm.Scope) error {
	if s.UpdatedAt.IsZero() {
		err := scope.SetColumn("UpdatedAt", time.Now())
		if err != nil {
			scope.Log("BeforeCreate error: %v", err)
		}
	}

	err := scope.SetColumn("CreatedAt", time.Now())
	if err != nil {
		scope.Log("BeforeCreate error: %v", err)
	}
	return nil
}

func (s *Subscription) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdatedAt", time.Now())
	if err != nil {
		scope.Log("BeforeUpdate error: %v", err)
	}
	return nil
}
