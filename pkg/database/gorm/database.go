package gorm

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // database driver
	"github.com/sumelms/microservice-course/pkg/config"
)

func Connect(cfg *config.Database) (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Database, cfg.Username, cfg.Password)

	db, err := gorm.Open(cfg.Driver, connString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the database")
	}

	return db, nil
}
