package skills

import "backend/src/entities"


type Reader interface {
	Read(dest interface{}, query string, args ...interface{}) error
}

func List(db Reader) ([]entities.Skill, error) {
	var sk []entities.Skill
	err := db.Read(&sk, "list-skills")
	return sk, err
}
