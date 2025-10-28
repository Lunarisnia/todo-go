package todo

import (
	"github.com/lunarisnia/todo-go/internal/server"
	"github.com/lunarisnia/todo-go/internal/todo/todoctl"
)

func Run() {
	s := server.NewServer()

	todoController := todoctl.NewToDoController()
	s.Register("/todo", todoController.CreateTask)
	s.Serve(":3000")
}
