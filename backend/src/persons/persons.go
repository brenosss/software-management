package persons

import (
	"backend/src/entities"
	"log"
)

type Reader interface {
	Read(dest interface{}, query string, args ...interface{}) error
}


func Get(db Reader, personID int64) entities.Person {
	person := entities.Person{}
	err := db.Read(&person, "get-person", personID)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Read(&person.PersonSkills, "get-skills", personID)
	if err != nil {
		log.Fatal(err)
	}
	return person

}
