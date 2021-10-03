package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-course/pkg/seed"
)

func Subscriptions() seed.Seed {
	return seed.Seed{
		Name: "CreateSubscriptions",
		Run: func(db *gorm.DB) error {
			s := Subscription{}
			return db.Create(s).Error
		},
	}
}
