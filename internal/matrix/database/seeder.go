package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-course/pkg/seed"
)

func Matrices() seed.Seed {
	return seed.Seed{
		Name: "CreateMatrices",
		Run: func(db *gorm.DB) error {
			u := &Matrix{}
			return db.Create(u).Error
		},
	}
}
