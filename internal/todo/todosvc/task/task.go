package task

type TaskStorage map[string]Task

type Task struct {
	TaskName string `json:"task_name"`
	Status   bool   `json:"status"`
}

type TaskRequest struct {
	TaskName string `json:"task_name"`
}
