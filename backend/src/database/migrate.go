package database

import (
	"log"
	"os"
	"time"

	"backend/src/files"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func fillsLanguageModel(db *sqlx.DB) {
	jsonLanguages := files.OpenJsonLanguages()
	stmt, _ := db.Prepare("INSERT INTO language(language_id, name, popularity, learning, created_at) values(?,?,?,?,?)")
	for _, element := range jsonLanguages {
		log.Println(element)
		stmt.Exec(uuid.NewString(), element.Name, element.Popularity, element.Learning, time.Now())
	}
}

func fillsSkillModel(db *sqlx.DB) {
	jsonSkills := files.OpenJsonSkills()
	stmt, _ := db.Prepare("INSERT INTO skill(skill_id, name, type, created_at) values(?,?,?,?)")
	for _, element := range jsonSkills {
		log.Println(element)
		stmt.Exec(uuid.NewString(), element.Name, element.Type, time.Now())
	}
}

func DBExists() bool {
	if _, err := os.Stat("software.db"); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Migrate() {
	db := GetConnection()
	log.Println(db)
	fillsLanguageModel(db)
	fillsSkillModel(db)
}
