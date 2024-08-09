package server

import (
	"gotodo/cmd/web"
	"gotodo/cmd/web/components"
	"gotodo/internal/todo"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func (s *FiberServer) RegisterFiberRoutes() {
	// s.App.Get("/", s.HelloWorldHandler)

	s.App.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(web.Files),
		PathPrefix: "assets",
		Browse:     false,
	}))

	s.App.Get("/", adaptor.HTTPHandler(templ.Handler(web.HelloForm(s.Todos))))
	s.App.Patch("/edit-todo/:id", s.EditTodo)
	s.App.Patch("/done-todo/:id", s.DoneTodo)
	s.App.Delete("/todo/:id", s.DeleteTodo)
	s.App.Post("/todo", s.AddTodo)

	s.App.Post("/hello", func(c *fiber.Ctx) error {
		return web.HelloWebHandler(c)
	})

}

func (s *FiberServer) EditTodo(c *fiber.Ctx) error {
	var id = c.Params("id")
	if s.Todos[id].Editing {
		s.Todos[id].Title = c.FormValue("todotitle")
	}
	s.Todos[id].Editing = !s.Todos[id].Editing

	return components.Todo(s.Todos[id]).
		Render(
			c.Context(),
			c.Response().BodyWriter(),
		)
}
func (s *FiberServer) DoneTodo(c *fiber.Ctx) error {
	var id = c.Params("id")
	s.Todos[id].Completed = !s.Todos[id].Completed

	return components.Todo(s.Todos[id]).
		Render(
			c.Context(),
			c.Response().BodyWriter(),
		)
}
func (s *FiberServer) DeleteTodo(c *fiber.Ctx) error {
	var id = c.Params("id")
	delete(s.Todos, id)
	return c.SendString("")
}
func (s *FiberServer) AddTodo(c *fiber.Ctx) error {
	var t = &todo.Todo{
		Id:        strconv.Itoa(len(s.Todos) + 1),
		Title:     c.FormValue("newtodo"),
		Completed: false,
		Editing:   false,
	}
	NewTodo(s, t)
	return components.Todo(t).
		Render(
			c.Context(),
			c.Response().BodyWriter(),
		)
}
