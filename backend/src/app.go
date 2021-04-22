package main

import (
	"flag"
	"log"
	"strconv"

	"backend/src/database"
	"backend/src/entities"
	"backend/src/languages"
	"backend/src/persons"
	"backend/src/projects"
	"backend/src/skills"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


const DatabaseName = "./software.db"

type Database interface {
	Query(dest interface{}, query string, args ...interface{}) error
	Run(query string, args ...interface{}) error
	Read(dest interface{}, query string, args ...interface{}) error
}

type Endpoints struct {
	db Database
}

func (e *Endpoints) ListLanguages(c *gin.Context) {
	ls, err := languages.List(e.db)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": ls,
	})
}

func (e *Endpoints) ListSkills(c *gin.Context) {
	sks, err := skills.List(e.db)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": sks,
	})
}

func (e *Endpoints) CreateRandomPerson(c *gin.Context) {
	p := entities.NewRandomPerson()
	err := persons.Create(e.db, &p)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": p,
	})
}

func (e *Endpoints) ListPerson(c *gin.Context) {
	ps, err := persons.List(e.db)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": ps,
	})
}

func (e *Endpoints) GetPerson(c *gin.Context) {
	personID, err := strconv.ParseInt(c.Param("person_id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	person, err := persons.Get(e.db, personID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": person,
	})
}

func (e *Endpoints) CreateRandomProject(c *gin.Context) {
	project := entities.NewRandomProject()
	err := projects.Create(e.db, &project)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": project,
	})
}

func CORSConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000/", "*"}
	return cors.New(config)
}


func main() {
	log.Println("Starting server...")
	dataMigrateFlag := flag.Bool("populate-database", false, "")
	flag.Parse()

	db, err := database.New(DatabaseName, "src/database/sql/queries")
	if err != nil {
		log.Panicln("Error while creating the database:", err)
	}
	endpoints := Endpoints{db: db}

	if dataMigrateFlag!= nil && *dataMigrateFlag {
		log.Println("Populating database...")
		err = db.Run("populate-languages")
		if err != nil {
			log.Fatalln("Error while populating languages:", err)
		}
		err = db.Run("populate-skills")
		if err != nil {
			log.Fatalln("Error while populating skill:", err)
		}
	}

	router := gin.Default()
	router.Use(CORSConfig())
	router.GET("/languages", endpoints.ListLanguages)
	router.GET("/skills", endpoints.ListSkills)

	//Person
	router.POST("/person", endpoints.CreateRandomPerson)
	router.GET("/person", endpoints.ListPerson)
	router.GET("/person/:person_id", endpoints.GetPerson)

	//Project
	router.POST("/project", endpoints.CreateRandomProject)

	log.Println("Server being initiated")
	err = router.Run()
	if err != nil {
		log.Panic(err)
	}
}
