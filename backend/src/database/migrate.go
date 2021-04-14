package database

import (
	"backend/src/files"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func fillsLanguageModel(db *sqlx.DB) {
	jsonLanguages := files.OpenJsonLanguages()
	stmt, _ := db.Prepare("INSERT INTO language(id, name, popularity, learning, created_at) values(?,?,?,?,?)")
	for _, element := range jsonLanguages {
		stmt.Exec(uuid.NewString(), element.Name, element.Popularity, element.Learning, time.Now())
	}
}

func fillsSkillModel(db *sqlx.DB) {
	jsonSkills := files.OpenJsonSkills()
	stmt, _ := db.Prepare("INSERT INTO skill(id, name, type, created_at) values(?,?,?,?)")
	for _, element := range jsonSkills {
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
	fmt.Print(db)
	fillsLanguageModel(db)
	fillsSkillModel(db)
}
