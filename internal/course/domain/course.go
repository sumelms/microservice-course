package domain

import (
	"github.com/jinzhu/gorm"
)

type UserID string

type Course struct {
	gorm.Model // @TODO Remove it from here
	UUID        string
	Title       string
	Subtitle    string
	Excerpt     string
	Description string
}
