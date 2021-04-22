package languages

import (
	"backend/src/entities"
)


type Reader interface {
	Read(dest interface{}, query string, args ...interface{}) error
}

func List(db Reader) ([]entities.Language, error) {
	var languages []entities.Language
	err := db.Read(&languages, "list-languages")
	return languages, err
}
