package domain

import (
	"time"

	"github.com/google/uuid"
)

// Subject struct
type Subject struct {
	ID        uint       `json:"id"`
	UUID      uuid.UUID  `json:"uuid"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Objective string     `json:"objective"`
	Credit    float32    `json:"credit"`
	Workload  float32    `json:"workload"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
