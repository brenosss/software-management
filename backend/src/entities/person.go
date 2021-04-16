package entities

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

type Person struct {
	PersonId     string `db:"person_id"`
	Name         string
	Email        string
	Phone        string
	Description  string
	Birthday     time.Time
	PersonSkills []PersonSkill `db:"person_skill"`
}

func NewRandonPerson() Person {
	return Person{
		PersonId:    uuid.NewString(),
		Name:        gofakeit.Name(),
		Email:       gofakeit.Email(),
		Phone:       gofakeit.Phone(),
		Description: gofakeit.Phrase(),
		Birthday:    gofakeit.DateRange(time.Now().AddDate(-60, 0, 0), time.Now().AddDate(-18, 0, 0)),
	}
}
