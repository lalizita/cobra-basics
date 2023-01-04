package todo

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Todos []TaskBody

type TaskBody struct {
	Id      string
	Task    string
	Checked bool
}

func (t *Todos) ReadTodos() error {
	content, err := ioutil.ReadFile("./internal/todo/todos.json")
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
