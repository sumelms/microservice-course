package domain

import "time"

type Matrix struct {
	ID          uint       `json:"id"`
	UUID        string     `json:"uuid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CourseID    string     `json:"course_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
