package database

import (
	"backend/src/entities"
	"log"
)

func GetLanguages() []entities.Language {
	db := GetConnection()
	languages := []entities.Language{}
	err := db.Select(&languages, "SELECT name, popularity, learning FROM language")
	if err != nil {
		log.Fatal(err)
	}
	return languages
}

func GetSkills() []entities.Skill {
	db := GetConnection()
	skills := []entities.Skill{}
	err := db.Select(&skills, "SELECT name, type FROM skills")
	if err != nil {
		log.Fatal(err)
	}
	return skills
}
