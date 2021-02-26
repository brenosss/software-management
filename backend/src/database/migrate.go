package database

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Migrate() {
	db := GetConnection()
  	db.AutoMigrate(&Product{})
}