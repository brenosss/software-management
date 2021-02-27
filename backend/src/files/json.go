package files

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"fmt"
)

type Language struct {
    Name string `json:"name"`
    Popularity string `json:"popularity"`
    Learning string    `json:"learning"`
    Use []string `json:"use"`
}

func OpenJson(fileName string) []Language{

    jsonFile, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
    }

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

	var languages []Language

    json.Unmarshal([]byte(byteValue), &languages)

    return languages
}