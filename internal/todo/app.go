package todo

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lunarisnia/todo-go/internal/middleware"
	"github.com/lunarisnia/todo-go/internal/server"
	"github.com/lunarisnia/todo-go/internal/todo/todoctl"
	"github.com/lunarisnia/todo-go/internal/todo/todosvc"
)

func Run() {
	s := server.NewServer()

	todoService := todosvc.NewToDoService()
	todoController := todoctl.NewToDoController(todoService)
	s.AddMiddleware(middleware.LogRequest)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	staticDir := filepath.Join(cwd, "static")
	fs := http.FileServer(http.Dir(staticDir))
	s.Get("/static/", http.StripPrefix("/static/", fs).ServeHTTP)

	s.Get("/todo", todoController.GetTasks)
	s.Post("/todo/create", todoController.CreateTask)
	s.Get("/todo/count", todoController.GetTaskCount)
	s.Serve(":3000")
}
