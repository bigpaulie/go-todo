package models

import (
	"log"

	"github.com/bigpaulie/go-todo/pkg/controllers/requests"
	"github.com/google/uuid"
)

type Todo struct {
	ID        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func (t *Todo) GenerateUUID() {
	t.ID = uuid.NewString()
}

// GetAllTodos fetch all todos from database and return items or error
func GetAllTodos() ([]Todo, error) {
	var todos []Todo

	query := `SELECT id, task, completed FROM todos ORDER BY id DESC`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Completed)
		if err != nil {
			log.Fatal(err)
			return todos, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

// CreateTodo create a new entry in the database or returns an error
func CreateTodo(task string) (Todo, error) {
	var todo Todo = Todo{}
	todo.GenerateUUID()

	todo.Task = task
	todo.Completed = false

	query := `INSERT INTO todos (id, task, completed) VALUES(?, ?, ?)`
	_, err := db.Exec(query, todo.ID, task, false)
	if err != nil {
		log.Fatal(err)
		return todo, err
	}

	return todo, nil
}

// FindOnById find item by id and return it or throw an error
func FindTodoById(id string) (Todo, error) {
	var todo Todo

	query := `SELECT id, task, completed FROM todos WHERE id = ?`
	row := db.QueryRow(query, id)

	err := row.Scan(&todo.ID, &todo.Task, &todo.Completed)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

// UpdateTodo update an item or return an error
func UpdateTodo(id string, request requests.UpdateTodoRequest) (Todo, error) {
	var todo Todo

	query := `UPDATE todos SET task = ?, completed = ? WHERE id = ?`
	_, err := db.Exec(query, request.Task, request.Completed, id)
	if err != nil {
		return todo, err
	}

	todo.ID = id
	todo.Task = request.Task
	todo.Completed = request.Completed

	return todo, nil
}

// DeleteTodo delete a specified item
func DeleteTodo(id string) error {
	query := `DELETE FROM todos WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
