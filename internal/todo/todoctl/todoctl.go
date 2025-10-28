package todoctl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lunarisnia/todo-go/internal/todo/todosvc"
	"github.com/lunarisnia/todo-go/internal/todo/todosvc/task"
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
	ctx := context.Background()

	var taskRequest task.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&taskRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := t.ToDoService.CreateTask(ctx, taskRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Created!")
}
