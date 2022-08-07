package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID         uint       `json:"id"`
	UUID       uuid.UUID  `json:"uuid"`
	UserID     uuid.UUID  `db:"user_id" json:"user_id"`
	CourseID   uuid.UUID  `db:"course_id" json:"course_id"`
	MatrixID   *uuid.UUID `db:"matrix_id" json:"matrix_id,omitempty"`
	ValidUntil *time.Time `db:"valid_until" json:"valid_until,omitempty"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deleted_at"`
}
