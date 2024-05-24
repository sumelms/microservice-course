package domain

import (
	"time"

	"github.com/google/uuid"
)

type MatrixFilters struct {
	CourseUUID uuid.UUID `json:"course_uuid,omitempty"`
}

type Matrix struct {
	UUID        uuid.UUID `db:"uuid"        json:"uuid"`
	Code        string    `db:"code"        json:"code"`
	Name        string    `db:"name"        json:"name"`
	Description string    `db:"description" json:"description"`
	CourseUUID  uuid.UUID `db:"course_uuid" json:"course_uuid"`
	CreatedAt   time.Time `db:"created_at"  json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"  json:"updated_at"`
}

type MatrixSubject struct {
	UUID        uuid.UUID `db:"uuid"         json:"uuid"`
	MatrixUUID  uuid.UUID `db:"matrix_uuid"  json:"matrix_uuid"`
	SubjectUUID uuid.UUID `db:"subject_uuid" json:"subject_uuid"`
	IsRequired  bool      `db:"is_required"  json:"is_required"`
	Group       string    `db:"group"        json:"group"`
	CreatedAt   time.Time `db:"created_at"   json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"   json:"updated_at"`
}

type DeletedMatrix struct {
	UUID      uuid.UUID `db:"uuid"       json:"uuid"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}

type DeletedMatrixSubject struct {
	UUID        uuid.UUID `db:"uuid"         json:"uuid"`
	SubjectUUID uuid.UUID `db:"subject_uuid" json:"subject_uuid"`
	MatrixUUID  uuid.UUID `db:"matrix_uuid"  json:"matrix_uuid"`
	DeletedAt   time.Time `db:"deleted_at"   json:"deleted_at"`
}
