package todosvc

import (
	"context"
	"errors"

	"github.com/lunarisnia/todo-go/internal/todo/todosvc/task"
)

type ToDoService interface {
	CreateTask(ctx context.Context, taskRequest task.TaskRequest) error
	GetTasks(ctx context.Context) ([]task.Task, error)
	GetTaskCount(ctx context.Context) int
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

func (t todoServiceImpl) GetTasks(ctx context.Context) ([]task.Task, error) {
	tasks := make([]task.Task, 0)
	for _, tsk := range t.TaskStorage {
		tasks = append(tasks, tsk)
	}

	if len(tasks) <= 0 {
		return nil, errors.New("no task")
	}

	return tasks, nil
}

func (t todoServiceImpl) GetTaskCount(ctx context.Context) int {
	return len(t.TaskStorage)
}
