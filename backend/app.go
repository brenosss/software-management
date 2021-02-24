package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
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

func main() {
	router := gin.Default()
	router.GET("/languages/info", LanguagesInfo)
	router.Run()
}