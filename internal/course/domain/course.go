package domain

import (
	"time"

	"github.com/google/uuid"
)

// Course struct.
type Course struct {
	ID          uint       `json:"id"`
	UUID        uuid.UUID  `json:"uuid"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Underline   string     `json:"underline"`
	Image       string     `json:"image"`
	ImageCover  string     `db:"image_cover"   json:"image_cover"`
	Excerpt     string     `json:"excerpt"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `db:"created_at"    json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"    json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"    json:"deleted_at"`
}
