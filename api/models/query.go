package models

import "gorm.io/gorm"

type (
	Query struct {
		gorm.Model
		Terms string
	}
)
