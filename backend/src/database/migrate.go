package database

import (
	"backend/src/database/models"
	"backend/src/files"
	"gorm.io/gorm"
	"os"
)

func fillsLanguageModel(db *gorm.DB) {
	jsonLanguages := files.OpenJsonLanguages()
	for _, element := range jsonLanguages {
		db.Create(&models.Language{
			Name: element.Name, Popularity: element.Popularity, Learning: element.Learning})
	}
}

func fillsSkillModel(db *gorm.DB) {
	jsonSkills := files.OpenJsonSkills()
	for _, element := range jsonSkills {
		db.Create(&models.Skill{
			Name: element.Name, Type: element.Type})
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
	if !DBExists(){
		db := GetConnection()
  		db.AutoMigrate(&models.Language{})
		db.AutoMigrate(&models.Skill{})
		fillsLanguageModel(db)
		fillsSkillModel(db)
	}
}