package database

import (
	"backend/src/entities"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
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
	err := db.Select(&skills, "SELECT skill_id, name, type FROM skill")
	if err != nil {
		log.Fatal(err)
	}
	return skills
}

func CreatePerson() entities.Person {
	createPersonQuery := `
		INSERT INTO person (
			person_id,
			name,
			email,
			phone,
			description,
			birthday,
			created_at
		) values(?,?,?,?,?,?,?)
	`
	db := GetConnection()
	person := entities.NewRandonPerson()
	stmt, err := db.Prepare(createPersonQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(person.PersonId, person.Name, person.Email, person.Phone, person.Description, person.Birthday, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	log.Print("New person created")
	skills := GetSkills()
	for _, skill := range skills {
		log.Print(skill)
		person.PersonSkills = append(person.PersonSkills, CreatePersonSkill(person, skill))
	}
	return person
}

func CreatePersonSkill(person entities.Person, skill entities.Skill) entities.PersonSkill {
	createPersonQuery := `
		INSERT INTO person_skill (
			person_skill_id,
			person_id,
			skill_id,
			value,
			progress,
			created_at
		) values (
			:person_skill_id,
			:person_id,
			:skill_id,
			:value,
			:progress,
			:created_at
		)
	`
	newPersonSkill := entities.PersonSkill{
		PersonSkillId: uuid.NewString(),
		PersonId:      person.PersonId,
		SkillId:       skill.SkillId,
		Value:         gofakeit.Number(0, 100),
		Progress:      gofakeit.Number(0, 100),
		CreatedAt:     time.Now(),
	}
	log.Print(newPersonSkill)
	db := GetConnection()
	_, err := db.NamedExec(createPersonQuery, newPersonSkill)
	if err != nil {
		log.Fatal(err)
	}
	return newPersonSkill
}
