package domain

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	UUID        uuid.UUID `json:"uuid"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Underline   string    `json:"underline"`
	Image       string    `json:"image,omitempty"`
	ImageCover  string    `db:"image_cover"             json:"image_cover,omitempty"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `db:"created_at"              json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"              json:"updated_at"`
}
