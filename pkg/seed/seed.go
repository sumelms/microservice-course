package seed

import (
	"github.com/jmoiron/sqlx"
)

type Seed struct {
	Name string
	Run  func(*sqlx.DB) error
}
