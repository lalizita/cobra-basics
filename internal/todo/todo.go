package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

const FILENAME = "./internal/todo/todos.json"

type Todos []TaskBody

type TaskBody struct {
	Task      string
	Checked   bool
	CreatedAt time.Time
}

func (t *Todos) ReadTodos() error {
	content, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return err
	}

	err = json.Unmarshal(content, t)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return err
	}

	return nil
}

func (t *Todos) Store() error {
	fmt.Println("data:", t)
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(FILENAME, data, 0644)
}
