package domain

import "time"

type Course struct {
	ID          uint       `json:"id"`
	UUID        string     `json:"uuid"`
	Title       string     `json:"title"`
	Subtitle    string     `json:"subtitle"`
	Excerpt     string     `json:"excerpt"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
