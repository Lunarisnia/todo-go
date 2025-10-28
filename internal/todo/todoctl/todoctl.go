package todoctl

import (
	"fmt"
	"net/http"

	"github.com/lunarisnia/todo-go/internal/todo/todosvc"
)

type ToDoController interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
}

type todoControllerImpl struct {
	ToDoService todosvc.ToDoService
}

func NewToDoController(todoService todosvc.ToDoService) ToDoController {
	return &todoControllerImpl{
		ToDoService: todoService,
	}
}

func (t todoControllerImpl) CreateTask(w http.ResponseWriter, r *http.Request) {
	err := t.ToDoService.CreateTask()
	if err != nil {
		fmt.Fprint(w, "Error!")
		return
	}
	fmt.Fprint(w, "Created!")
}
