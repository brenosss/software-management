package entities

import "time"

type PersonSkill struct {
	ID        string `db:"person_skill_id"`
	PersonID  string `db:"person_id"`
	SkillID   string `db:"skill_id"`
	Value     int       `db:"value"`
	Progress  int       `db:"progress"`
	CreatedAt time.Time `db:"created_at"`

	Person    Person
	Skill     Skill     `db:"skill"`
}
