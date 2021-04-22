package tasks

import "backend/src/entities"


type Writer interface {
	Query(dest interface{}, query string, args ...interface{}) error
}


func Create(db Writer, task *entities.Task) error {
	return db.Query(&task.ID, "create-task",
		task.ProjectID, task.Name, task.Type, task.Description, task.Complexity, task.Duration, task.Progress)
}
