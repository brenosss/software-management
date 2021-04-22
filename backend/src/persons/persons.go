package persons

import "backend/src/entities"


type Reader interface {
	Read(dest interface{}, query string, args ...interface{}) error
}

type Writer interface {
	Query(dest interface{}, query string, args ...interface{}) error
}


func Get(db Reader, personID int64) (entities.Person, error) {
	person := entities.Person{}
	err := db.Read(&person, "get-person", personID)
	if err != nil {
		return person, err
	}
	err = db.Read(&person.PersonSkills, "get-skills", personID)
	return person, err

}

func List(db Reader) ([]entities.Person, error) {
	var persons []entities.Person
	err := db.Read(&persons, "list-persons")
	return persons, err
}

func Create(db Writer, person *entities.Person) error {
	// TODO put it inside a transaction
	err := db.Query(&person.ID, "create-person",
		person.ID, person.Name, person.Email, person.Phone, person.Description, person.Birthday)
	if err != nil {
		return err
	}
	for _, ps := range person.PersonSkills {
		AssignSkill(db, &ps)
	}
	return nil
}

func AssignSkill(db Writer, personSkill *entities.PersonSkill) error {
	return db.Query(&personSkill.ID, "assign-skill",
		personSkill.PersonID, personSkill.SkillID, personSkill.Value, personSkill.Progress)
}
