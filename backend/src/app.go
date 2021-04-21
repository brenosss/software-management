package main

import (
	"backend/src/database"
	"backend/src/database/queries"
	"backend/src/persons"
	"flag"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

func PersonGet(context *gin.Context) {
	// TODO handle errors
	personID, _ := strconv.ParseInt(context.Param("person_id"), 10, 64)
	db, _ := database.New(database.DatabaseName, "backend/src/database/sql/queries")
	person := persons.Get(db, personID)
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
	dataMigrateFlag := flag.Bool("datamigrate", false, "")
	flag.Parse()
	if *dataMigrateFlag {
		database.Migrate()
		return
	}
	router := gin.Default()
	router.Use(CORSConfig())
	router.GET("/languages", LanguagesInfo)
	router.GET("/skills", SkillsInfo)

	//Person
	router.POST("/person", PersonCreate)
	router.GET("/person", PersonList)
	router.GET("/person/:person_id", PersonGet)

	//Project
	router.POST("/project", ProjectCreate)

	router.Run()
}
