package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"backend/src/database"
)

func OpenJson(fileName string) []map[string]interface{}{

    jsonFile, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
    }

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var result []map[string]interface{}

    json.Unmarshal([]byte(byteValue), &result)

    return result
}

func LanguagesInfo(context *gin.Context) {
	content := OpenJson("languages_definition.json")
	context.JSON(200, gin.H{
		"data": content,
	})
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
	router.GET("/languages/info", LanguagesInfo)
	router.Run()
}