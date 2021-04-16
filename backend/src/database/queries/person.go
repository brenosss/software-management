package queries

import (
	"backend/src/database"
	"backend/src/entities"
	"log"
)

func GetPeople() []entities.Person {
	db := database.GetConnection()
	people := []entities.Person{}
	err := db.Select(&people, "SELECT person_id, name, email FROM person")
	if err != nil {
		log.Fatal(err)
	}
	return people
}

func GetPerson(personId string) entities.Person {
	db := database.GetConnection()
	personQuery := `
		SELECT
			person_id, name, email, phone, description
		FROM person
		WHERE person.person_id=$1
	`
	personSkillQuery := `
		SELECT
			person_skill_id,
			value,
			progress,
			skill.skill_id "skill.skill_id",
			skill.name "skill.name",
			skill.type "skill.type"
		FROM person_skill
		JOIN skill ON skill.skill_id = person_skill.skill_id
		WHERE person_skill.person_id=$1
	`
	person := entities.Person{}
	err := db.Get(&person, personQuery, personId)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Select(&person.PersonSkills, personSkillQuery, personId)
	if err != nil {
		log.Fatal(err)
	}
	return person
}
