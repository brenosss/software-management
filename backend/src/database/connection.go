package database

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

func GetConnection() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("software.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	return db
}
