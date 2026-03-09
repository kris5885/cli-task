package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Task represents a single task item.
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

// TaskManager handles storage of tasks.
type TaskManager struct {
	FilePath string
}

// NewTaskManager creates a new task manager with the default file path.
func NewTaskManager() (*TaskManager, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(home, ".tasks.json")
	return &TaskManager{FilePath: filePath}, nil
}

// LoadTasks reads tasks from the JSON file.
func (tm *TaskManager) LoadTasks() ([]Task, error) {
	if _, err := os.Stat(tm.FilePath); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(tm.FilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	var tasks []Task
	if len(data) == 0 {
		return tasks, nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("could not parse json: %v", err)
	}

	return tasks, nil
}

// SaveTasks writes tasks to the JSON file.
func (tm *TaskManager) SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal json: %v", err)
	}

	return os.WriteFile(tm.FilePath, data, 0644)
}

// AddTask adds a new task and saves it.
func (tm *TaskManager) AddTask(description string) (*Task, error) {
	tasks, err := tm.LoadTasks()
	if err != nil {
		return nil, err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	task := Task{
		ID:          id,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, task)
	if err := tm.SaveTasks(tasks); err != nil {
		return nil, err
	}

	return &task, nil
}

// DoTask marks a task as completed.
func (tm *TaskManager) DoTask(id int) error {
	tasks, err := tm.LoadTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return tm.SaveTasks(tasks)
}

// DeleteTask removes a task.
func (tm *TaskManager) DeleteTask(id int) error {
	tasks, err := tm.LoadTasks()
	if err != nil {
		return err
	}

	var updated []Task
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		updated = append(updated, t)
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return tm.SaveTasks(updated)
}
