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

type Skill struct {
    Name string `json:"name"`
    Type string `json:"type"`
}

func OpenJson(fileName string) *os.File{

    jsonFile, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
    }
    return jsonFile
}

func OpenJsonLanguages() []Language{
    jsonFile := OpenJson("languages_definition.json")
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)

	var languages []Language

    json.Unmarshal([]byte(byteValue), &languages)

    return languages
}

func OpenJsonSkills() []Skill{
    jsonFile := OpenJson("skills_definition.json")
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)

	var skills []Skill

    json.Unmarshal([]byte(byteValue), &skills)

    return skills
}