package server

import (
	"gotodo/internal/todo"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
	Todos map[string]*todo.Todo
}

func NewTodo(s *FiberServer, newTodo *todo.Todo) {
	s.Todos[newTodo.Id] = newTodo
}

func New() *FiberServer {
	var todos = map[string]*todo.Todo{}

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "gotodo",
			AppName:      "gotodo",
		}),
		Todos: todos,
	}
	NewTodo(server, &todo.Todo{
		Id:        "xxxx",
		Title:     "sample todo",
		Completed: false,
		Editing:   false,
	})
	return server
}
