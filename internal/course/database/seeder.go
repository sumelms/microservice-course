package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-course/pkg/seed"
)

func Courses() seed.Seed {
	return seed.Seed{
		Name: "CreateCourses",
		Run: func(db *gorm.DB) error {
			u := &Course{}
			return db.Create(u).Error
		},
	}
}
