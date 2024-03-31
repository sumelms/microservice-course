package domain

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionFilters struct {
	CourseUUID uuid.UUID `json:"course_uuid,omitempty"`
	UserUUID   uuid.UUID `json:"user_uuid,omitempty"`
}

type Course struct {
	UUID uuid.UUID `db:"uuid" json:"uuid"`
	Code string    `db:"code" json:"code"`
	Name string    `db:"name" json:"name"`
}

type Matrix struct {
	UUID *uuid.UUID `db:"uuid" json:"uuid"`
	Code *string    `db:"code" json:"code"`
	Name *string    `db:"name" json:"name"`
}

type Subscription struct {
	UUID       uuid.UUID  `db:"uuid"        json:"uuid"`
	UserUUID   uuid.UUID  `db:"user_uuid"   json:"user_uuid,omitempty"`
	Course     *Course    `db:"courses"     json:"course,omitempty"`
	CourseUUID *uuid.UUID `db:"course_uuid" json:"course_uuid,omitempty"`
	Matrix     *Matrix    `db:"matrices"    json:"matrix,omitempty"`
	MatrixUUID *uuid.UUID `db:"matrix_uuid" json:"matrix_uuid,omitempty"`
	Role       string     `db:"role"        json:"role,omitempty"`
	Reason     *string    `db:"reason"      json:"reason,omitempty"`
	ExpiresAt  *time.Time `db:"expires_at"  json:"expires_at,omitempty"`
	CreatedAt  time.Time  `db:"created_at"  json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"  json:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"  json:"deleted_at,omitempty"`
}
