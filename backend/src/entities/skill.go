package entities


type Skill struct {
	ID   int64 `db:"skill_id"`
	Name string
	Type string
}
