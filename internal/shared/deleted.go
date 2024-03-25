package shared

import (
	"time"

	"github.com/google/uuid"
)

type Deleted struct {
	UUID      uuid.UUID `json:"uuid"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}
