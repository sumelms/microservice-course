package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type Course struct {
	UUID uuid.UUID `db:"course.uuid" json:"uuid"`
	Name string    `db:"course.name" json:"name"`
}

type Matrix struct {
	UUID uuid.UUID `db:"matrix.uuid" json:"uuid"`
	Name string    `db:"matrix.name" json:"name"`
}

type Subscription struct {
	UUID       uuid.UUID  `db:"uuid"        json:"uuid"`
	User       User       `db:"-"           json:"user,omitempty"`
	UserUUID   uuid.UUID  `db:"user_uuid"   json:"user_uuid"`
	Course     Course     `db:"-"           json:"course,omitempty"`
	CourseUUID uuid.UUID  `db:"course_uuid" json:"course_uuid"`
	Matrix     Matrix     `db:"-"           json:"matrix,omitempty"`
	MatrixUUID *uuid.UUID `db:"matrix_uuid" json:"matrix_uuid,omitempty"`
	Role       string     `db:"role"        json:"role"`
	Reason     *string    `db:"reason"      json:"reason,omitempty"`
	ExpiresAt  *time.Time `db:"expires_at"  json:"expires_at,omitempty"`
	CreatedAt  time.Time  `db:"created_at"  json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"  json:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"  json:"deleted_at,omitempty"`
}
