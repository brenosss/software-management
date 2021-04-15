package entities

import "time"

type PersonSkill struct {
	PersonSkillId string `db:"person_skill_id"`
	PersonId      string `db:"person_id"`
	SkillId       string `db:"skill_id"`
	Person        Person
	Skill         Skill
	Value         int
	Progress      int
	CreatedAt     time.Time `db:"created_at"`
}
