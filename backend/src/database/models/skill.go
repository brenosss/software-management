package models

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Name  string
	Type string
}