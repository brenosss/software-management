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

func NewRandomPerson(skills []Skill) Person {
	p := Person{
		ID:          gofakeit.Int64(),
		Name:        gofakeit.Name(),
		Email:       gofakeit.Email(),
		Phone:       gofakeit.Phone(),
		Description: gofakeit.Phrase(),
		Birthday:    gofakeit.DateRange(time.Now().AddDate(-60, 0, 0), time.Now().AddDate(-18, 0, 0)),
	}
	for _, s := range skills {
		p.PersonSkills = append(p.PersonSkills, PersonSkill{
			PersonID:      p.ID,
			SkillID:       s.ID,
			Value:         gofakeit.Number(0, 100),
			Progress:      gofakeit.Number(0, 100),
		})
	}
	return p
}
