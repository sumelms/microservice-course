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
	db.AutoMigrate(&Matrix{})

	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r Repository) Create() {
	panic("implement me")
}

func (r Repository) Find() {
	panic("implement me")
}

func (r Repository) Update() {
	panic("implement me")
}

func (r Repository) Delete() {
	panic("implement me")
}

func (r Repository) List() {
	panic("implement me")
}
