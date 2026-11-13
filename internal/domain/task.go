package domain

import (
	"errors"
	"time"
)

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)

var validStatuses = map[TaskStatus]bool{
	TaskStatusPending:    true,
	TaskStatusInProgress: true,
	TaskStatusCompleted:  true,
}

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewTask(id, title, description string, status TaskStatus) (*Task, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if description == "" {
		return nil, errors.New("description cannot be empty")
	}
	if !validStatuses[status] {
		return nil, errors.New("invalid status")
	}

	now := time.Now()
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (t *Task) UpdateTitle(title string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	t.Title = title
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) UpdateDescription(description string) error {
	if description == "" {
		return errors.New("description cannot be empty")
	}
	t.Description = description
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) UpdateStatus(status TaskStatus) error {
	if !validStatuses[status] {
		return errors.New("invalid status")
	}
	t.Status = status
	t.UpdatedAt = time.Now()
	return nil
}