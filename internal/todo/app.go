package todo

import "github.com/lunarisnia/todo-go/internal/server"

func Run() {
	s := server.NewServer()
	s.Serve(":3000")
}
