package entities

import "time"

type PersonSkill struct {
	PersonSkillId string `db:"person_skill_id"`
	PersonId      string `db:"person_id"`
	SkillId       string `db:"skill_id"`
	Person        Person
	Skill         Skill     `db:"skill"`
	Value         int       `db:"value"`
	Progress      int       `db:"progress"`
	CreatedAt     time.Time `db:"created_at"`
}
