package entities

type Skill struct {
	SkillId string `db:"skill_id"`
	Name    string
	Type    string
}
