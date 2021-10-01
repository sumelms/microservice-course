package database

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&Subscription{})

	return &Repository{db: db, logger: logger}
}
