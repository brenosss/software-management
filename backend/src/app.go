package main

import (
	"backend/src/database"
	"backend/src/database/queries"
	"flag"

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
	personId := context.Param("person_id")
	person := queries.GetPerson(personId)
	context.JSON(200, gin.H{
		"data": person})
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

	router.Run()
}
