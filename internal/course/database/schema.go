package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Course struct
type Course struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string    `gorm:"size:100"`
	Subtitle    string    `gorm:"size:100"`
	Excerpt     string    `gorm:"size:144"`
	Description string    `gorm:"size:255"`
}

func (p *Course) BeforeCreate(scope *gorm.Scope) error {
	if p.UpdatedAt.IsZero() {
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

func (p *Course) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdatedAt", time.Now())
	if err != nil {
		scope.Log("BeforeUpdate error: %v", err)
	}
	return nil
}
