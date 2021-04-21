package entities

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

type Project struct {
	ProjectId string `db:"project_id"`
	Name      string `db:"name"`
	Tasks     []Task `db:"task"`
}

func NewRandonProject() Project {
	projectId := uuid.NewString()
	tasks := []Task{}
	for i := 1; i <= 100; i++ {
		tasks = append(tasks, NewRandonTask(projectId))
	}
	return Project{
		ProjectId: projectId,
		Name:      gofakeit.AppName(),
		Tasks:     tasks,
	}
}
