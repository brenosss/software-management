package entities

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Task struct {
	ID          int64     `db:"task_id"`
	ProjectID   int64    `db:"project_id"`
	PersonID    int64    `db:"person_id"`
	Name        string    `db:"name"`
	Type        string    `db:"type"`
	Description string    `db:"description"`
	Complexity  int       `db:"complexity"`
	Duration    int       `db:"duration"`
	Progress    int       `db:"progress"`
	CreatedAt   time.Time `db:"created_at"`

	Person  Person `db:"person"`
	Project Skill  `db:"project"`
}

func NewRandomTask(projectID int64) Task {
	return Task{
		ID:          gofakeit.Int64(),
		ProjectID:   projectID,
		Name:        gofakeit.HackerVerb() + " " + gofakeit.HackerNoun(),
		Type:        gofakeit.RandomString([]string{"BUG", "FEATURE", "REFACTORING"}),
		Description: gofakeit.HackerPhrase(),
		Complexity:  gofakeit.Number(1, 5),
		Duration:    gofakeit.Number(1, 5),
		Progress:    0,
	}
}
