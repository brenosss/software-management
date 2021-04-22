package projects

import (
	"backend/src/entities"
	"backend/src/tasks"
)

type Reader interface {
	Read(dest interface{}, query string, args ...interface{}) error
}

type Writer interface {
	Query(dest interface{}, query string, args ...interface{}) error
}

func Create(db Writer, project *entities.Project) error {
	err := db.Query(project.ID, "create-project", project.Name)
	for _, task := range project.Tasks {
		tasks.Create(db, &task)
	}
	return err
}
