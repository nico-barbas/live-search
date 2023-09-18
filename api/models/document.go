package models

import "gorm.io/gorm"

type (
	Document struct {
		gorm.Model
		Name    string
		Content string
	}
)
