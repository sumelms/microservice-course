package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type Matrix struct {
	gorm.Model
	UUID        uuid.UUID       `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title       string          `gorm:"size:100"`
	Description string          `gorm:"size:255"`
	CourseID    domain.CourseID `gorm:"index;"`
}

func (m *Matrix) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	scope.SetColumn("UUID", id.String())

	if m.UpdatedAt.IsZero() {
		err := scope.SetColumn("UpdatedAt", time.Now())
		if err != nil {
			scope.Log("BeforeCreate error: %v", err)
		}
	}

	err = scope.SetColumn("CreatedAt", time.Now())
	if err != nil {
		scope.Log("BeforeCreate error: %v", err)
	}
	return nil
}

func (m *Matrix) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdatedAt", time.Now())
	if err != nil {
		scope.Log("BeforeUpdate error: %v", err)
	}
	return nil
}
