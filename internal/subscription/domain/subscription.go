package domain

import "time"

type Subscription struct {
	ID         uint       `json:"id"`
	UserID     string     `json:"user_id"`
	CourseID   string     `json:"course_id"`
	MatrixID   string     `json:"matrix_id"`
	ValidUntil *time.Time `json:"valid_until"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
