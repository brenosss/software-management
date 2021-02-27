package models

import (
	"gorm.io/gorm"
)

type Language struct {
	gorm.Model
	Name  string
	Popularity string
	Learning string
}