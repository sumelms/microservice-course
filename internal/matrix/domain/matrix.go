package domain

import "time"

type Matrix struct {
	ID          uint
	UUID        string
	Title       string
	Description string
	CourseID    string `json:"course_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
