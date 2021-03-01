package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"backend/src/database"
	"backend/src/database/models"
)

func LanguagesInfo(context *gin.Context) {
	db := database.GetConnection()
	var languages []models.Language
	db.Find(&languages)
	context.JSON(200, gin.H{
		"data": languages})
}

func SkillsInfo(context *gin.Context) {
	db := database.GetConnection()
	var skills []models.Skill
	db.Find(&skills)
	context.JSON(200, gin.H{
		"data": skills})
}

func CORSConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000/", "*"}
	return cors.New(config)
}

func main() {
	database.Migrate()
	router := gin.Default()
	router.Use(CORSConfig())
	router.GET("/languages", LanguagesInfo)
	router.GET("/skills", SkillsInfo)
	router.Run()
}