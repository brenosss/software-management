package entities

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

type Task struct {
	TaskId    string `db:"task_id"`
	ProjectId string `db:"project_id"`
	PersonId  string `db:"person_id"`

	Person  Person `db:"person"`
	Project Skill  `db:"project"`

	Name        string    `db:"name"`
	Type        string    `db:"type"`
	Description string    `db:"description"`
	Complexity  int       `db:"complexity"`
	Duration    int       `db:"duration"`
	Progress    int       `db:"progress"`
	CreatedAt   time.Time `db:"created_at"`
}

func NewRandonTask(projectId string) Task {
	return Task{
		TaskId:      uuid.NewString(),
		ProjectId:   projectId,
		Name:        gofakeit.HackerVerb() + " " + gofakeit.HackerNoun(),
		Type:        gofakeit.RandomString([]string{"BUG", "FEATURE", "REFACTORING"}),
		Description: gofakeit.HackerPhrase(),
		Complexity:  gofakeit.Number(1, 5),
		Duration:    gofakeit.Number(1, 5),
		Progress:    0,
	}
}
