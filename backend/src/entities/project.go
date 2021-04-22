package entities

import (
	"github.com/brianvoe/gofakeit/v6"
)

type Project struct {
	ID    int64 `db:"project_id"`
	Name  string `db:"name"`
	Tasks []Task `db:"task"`
}

func NewRandomProject() Project {
	projectID := gofakeit.Int64()
	var tasks []Task
	for i := 1; i <= 100; i++ {
		tasks = append(tasks, NewRandomTask(projectID))
	}
	return Project{
		ID:    projectID,
		Name:  gofakeit.AppName(),
		Tasks: tasks,
	}
}
