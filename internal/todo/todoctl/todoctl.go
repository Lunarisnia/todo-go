package todoctl

import (
	"fmt"
	"net/http"
)

type ToDoController interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
}

type todoControllerImpl struct {
}

func NewToDoController() ToDoController {
	return &todoControllerImpl{}
}

func (t todoControllerImpl) CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Created!")
}
