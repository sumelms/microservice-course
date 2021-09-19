package domain

import (
	"github.com/jinzhu/gorm"
)

type UserID string

type Course struct {
	gorm.Model
	UUID        string
	Title       string
	Subtitle    string
	Excerpt     string
	Description string
}
