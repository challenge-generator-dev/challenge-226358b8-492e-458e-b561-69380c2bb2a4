package application

import (
	"context"
	"errors"

	"github.com/user/project/internal/domain"
)

type TaskRepository interface {
	Save(ctx context.Context, task *domain.Task) error
	FindByID(ctx context.Context, id string) (*domain.Task, error)
	FindAll(ctx context.Context) ([]*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, title, description string, status domain.TaskStatus) (*domain.Task, error) {
	task, err := domain.NewTask("", title, description, status)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(ctx, task); err != nil {
		return nil, errors.New("failed to save task")
	}

	return task, nil
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, errors.New("failed to list tasks")
	}
	return tasks, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id, title, description string, status domain.TaskStatus) (*domain.Task, error) {
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("task not found")
	}

	if title != "" {
		if err := task.UpdateTitle(title); err != nil {
			return nil, err
		}
	}

	if description != "" {
		if err := task.UpdateDescription(description); err != nil {
			return nil, err
		}
	}

	if status != "" {
		if err := task.UpdateStatus(status); err != nil {
			return nil, err
		}
	}

	if err := s.repo.Update(ctx, task); err != nil {
		return nil, errors.New("failed to update task")
	}

	return task, nil
}