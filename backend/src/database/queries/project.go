package queries

import (
	"backend/src/database"
	"backend/src/entities"
	"log"
)

func CreateProject() entities.Project {
	createProjectQuery := `
		INSERT INTO project (
			project_Id,
			name
		) values(?,?)
	`
	db := database.GetConnection()
	project := entities.NewRandonProject()
	stmt, err := db.Prepare(createProjectQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(project.ProjectId, project.Name)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("New person project")
	for _, task := range project.Tasks {
		log.Print(task)
		CreateTask(task)
	}
	return project
}

func CreateTask(task entities.Task) entities.Task {
	createTaskQuery := `
	INSERT INTO task (
		task_id,
		project_id,
		name,
		type,
		description,
		complexity,
		duration,
		progress
	) values(
		:task_id,
		:project_id,
		:name,
		:type,
		:description,
		:complexity,
		:duration,
		:progress
	)
	`
	db := database.GetConnection()
	_, err := db.NamedExec(createTaskQuery, task)
	if err != nil {
		log.Fatal(err)
	}
	return task
}
