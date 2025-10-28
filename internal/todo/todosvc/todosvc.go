package todosvc

import (
	"context"
	"errors"

	"github.com/lunarisnia/todo-go/internal/todo/todosvc/task"
)

type ToDoService interface {
	CreateTask(ctx context.Context, taskRequest task.TaskRequest) error
}

type todoServiceImpl struct {
	TaskStorage task.TaskStorage
}

func NewToDoService() ToDoService {
	return &todoServiceImpl{
		TaskStorage: make(task.TaskStorage),
	}
}

func (t todoServiceImpl) CreateTask(ctx context.Context, taskRequest task.TaskRequest) error {
	if _, exist := t.TaskStorage[taskRequest.TaskName]; exist {
		return errors.New("task already existed")
	}
	t.TaskStorage[taskRequest.TaskName] = task.Task{
		TaskName: taskRequest.TaskName,
		Status:   false,
	}
	return nil
}
