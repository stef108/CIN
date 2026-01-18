package models

import "time"

// TaskStatus represents the state of a task
type TaskStatus string

const (
	StatusTodo       TaskStatus = "TODO"
	StatusInProgress TaskStatus = "IN_PROGRESS"
	StatusDone       TaskStatus = "DONE"
)

// Task represents a unit of work
type Task struct {
	ID        int
	Title     string
	Status    TaskStatus
	CreatedAt time.Time
}

// NewTask creates a task with defaults
func NewTask(id int, title string) Task {
	return Task{
		ID:        id,
		Title:     title,
		Status:    StatusTodo,
		CreatedAt: time.Now(),
	}
}
