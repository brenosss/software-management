package entities

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Person struct {
	ID           int64 `db:"person_id"`
	Name         string
	Email        string
	Phone        string
	Description  string
	Birthday     time.Time
	PersonSkills []PersonSkill `db:"person_skill"`
}

func NewRandomPerson() Person {
	return Person{
		ID:          gofakeit.Int64(),
		Name:        gofakeit.Name(),
		Email:       gofakeit.Email(),
		Phone:       gofakeit.Phone(),
		Description: gofakeit.Phrase(),
		Birthday:    gofakeit.DateRange(time.Now().AddDate(-60, 0, 0), time.Now().AddDate(-18, 0, 0)),
	}
}
