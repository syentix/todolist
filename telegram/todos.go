package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/aodin/date"
)

type toDo struct {
	id        int    `json:"id"`
	todo      string `json:"todo"`
	date      string `json:"date"`
	completed bool   `json:"completed"`
}

func AddToDo(data Package) {

	// Getting the Path.
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	path := pwd + "/telegram/storage/" + data.userID + ".json"

	// Check if user has used bot before. If not create storage file.
	jsonFile, err := os.Open(path)
	if err != nil {
		os.Create(path)
	}
	defer jsonFile.Close()

	// Get the current JSON Data
	jsonData, _ := ioutil.ReadAll(jsonFile)

	var curJSON []toDo
	err = json.Unmarshal(jsonData, &curJSON)

	// Create new Struct and fill with data.
	curDate := date.Today().String()
	var thisID int = 0
	if len(curJSON) != 0 {
		thisID = getID(curJSON)
	}

	var newToDo toDo
	newToDo.completed = false
	newToDo.id = thisID
	newToDo.todo = data.text
	newToDo.date = curDate

	curJSON = append(curJSON, newToDo)

	result, _ := json.MarshalIndent(&curJSON, "", " ")
	s := string(result[:])
	fmt.Println(s)
	_ = ioutil.WriteFile(path, result, 0644)

}

func getID(todos []toDo) int {
	c := rand.Intn(255-0) + 0
	for _, todo := range todos {
		if todo.id == c {
			getID(todos)
		}
	}
	return c
}
