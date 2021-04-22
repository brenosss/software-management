package main

import (
	"flag"
	"log"
	"strconv"

	"backend/src/database"
	"backend/src/database/queries"
	"backend/src/persons"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Database interface {
	Query(dest interface{}, query string, args ...interface{}) error
	Run(query string, args ...interface{}) error
	Read(dest interface{}, query string, args ...interface{}) error
}

type Endpoints struct {
	db Database
}

func LanguagesInfo(context *gin.Context) {
	languages := database.GetLanguages()
	context.JSON(200, gin.H{
		"data": languages})
}

func SkillsInfo(context *gin.Context) {
	skills := database.GetSkills()
	context.JSON(200, gin.H{
		"data": skills})
}

func PersonCreate(context *gin.Context) {
	person := database.CreatePerson()
	context.JSON(200, gin.H{
		"data": person})
}

func PersonList(context *gin.Context) {
	people := queries.GetPeople()
	context.JSON(200, gin.H{
		"data": people})
}

func (e *Endpoints) PersonGet(context *gin.Context) {
	// TODO handle errors
	personID, _ := strconv.ParseInt(context.Param("person_id"), 10, 64)
	person := persons.Get(e.db, personID)
	context.JSON(200, gin.H{
		"data": person})
}

func ProjectCreate(context *gin.Context) {
	project := queries.CreateProject()
	context.JSON(200, gin.H{
		"data": project})
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

	db, err := database.New(database.DatabaseName, "src/database/sql/queries")
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
	router.GET("/languages", LanguagesInfo)
	router.GET("/skills", SkillsInfo)

	//Person
	router.POST("/person", PersonCreate)
	router.GET("/person", PersonList)
	router.GET("/person/:person_id", endpoints.PersonGet)

	//Project
	router.POST("/project", ProjectCreate)

	log.Println("Server being initiated")
	err = router.Run()
	if err != nil {
		log.Panic(err)
	}
}
