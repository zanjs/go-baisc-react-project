package repo

import (
	"fmt"

	"../model"
)

var currentID int

var todos model.Todos

// Give us some seed data
func init() {
	RepoCreateTodo(model.Todo{Name: "Write presentation"})
	RepoCreateTodo(model.Todo{Name: "Host meetup"})
}

func GetTodo() model.Todos {
	return todos
}

//RepoFindTodo will find a Todo item by the id
func RepoFindTodo(id int) model.Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	// return empty Todo if not found
	return model.Todo{}
}

//RepoCreateTodo will create a Todo item
func RepoCreateTodo(t model.Todo) model.Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

//RepoDestroyTodo will destroy all Todo items in the list
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
