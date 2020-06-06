package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreateAt  time.Time `json:"create_at"`
}

var todoMap map[int]*Todo

func GetTodos() []*Todo {
	return nil
}

func AddTodo(name string) *Todo {
	return nil
}

func RemoveTodo(id int) bool {
	return false
}

func CompleteTodo(id int, complete bool) bool {
	return false
}
