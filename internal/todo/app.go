package todo

import (
	"github.com/lunarisnia/todo-go/internal/server"
	"github.com/lunarisnia/todo-go/internal/todo/todoctl"
	"github.com/lunarisnia/todo-go/internal/todo/todosvc"
)

func Run() {
	s := server.NewServer()

	todoService := todosvc.NewToDoService()
	todoController := todoctl.NewToDoController(todoService)
	s.Post("/todo", todoController.CreateTask)
	s.Serve(":3000")
}
