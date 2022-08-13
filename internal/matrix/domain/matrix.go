package domain

import (
	"time"

	"github.com/google/uuid"
)

type Matrix struct {
	ID          uint       `json:"id"`
	UUID        uuid.UUID  `json:"uuid"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CourseID    uuid.UUID  `db:"course_id" json:"course_id"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}

type MatrixSubject struct {
	ID         uint      `json:"id"`
	SubjectID  uuid.UUID `db:"subject_id" json:"subject_id"`
	MatrixID   uuid.UUID `db:"matrix_id" json:"matrix_id"`
	IsRequired bool      `db:"is_required" json:"is_required"`
}
