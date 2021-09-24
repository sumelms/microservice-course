package domain

import "time"

type CourseID uint

type Matrix struct {
	ID          uint
	UUID        string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	CourseID    CourseID
}
