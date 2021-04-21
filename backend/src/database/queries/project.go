package queries

import (
	"backend/src/entities"
)

func CreateProject() entities.Project {
	return entities.NewRandonProject()
}
