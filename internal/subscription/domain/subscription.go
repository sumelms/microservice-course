package domain

import "time"

type CourseID string
type UserID string
type MatrixID string

type Subscription struct {
	ID         uint
	UserID     UserID
	CourseID   CourseID
	MatrixID   MatrixID
	ValidUntil *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
