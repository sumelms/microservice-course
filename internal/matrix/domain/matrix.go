package domain

import (
	"time"

	"github.com/google/uuid"
)

type MatrixFilters struct {
	CourseUUID uuid.UUID `json:"course_uuid,omitempty"`
}

type Course struct {
	UUID uuid.UUID `db:"uuid" json:"uuid"`
	Code string    `db:"code" json:"code"`
	Name string    `db:"name" json:"name"`
}

type Matrix struct {
	UUID        uuid.UUID  `db:"uuid"        json:"uuid"`
	Code        string     `db:"code"        json:"code"`
	Name        string     `db:"name"        json:"name"`
	Description string     `db:"description" json:"description"`
	CourseUUID  *uuid.UUID `db:"course_uuid" json:"course_uuid,omitempty"`
	Course      *Course    `db:"courses"     json:"course,omitempty"`
	CreatedAt   time.Time  `db:"created_at"  json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"  json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"  json:"deleted_at,omitempty"`
}

type MatrixSubject struct {
	ID         uint      `json:"id"`
	SubjectID  uuid.UUID `db:"subject_id"  json:"subject_id"`
	MatrixID   uuid.UUID `db:"matrix_id"   json:"matrix_id"`
	IsRequired bool      `db:"is_required" json:"is_required"`
}
