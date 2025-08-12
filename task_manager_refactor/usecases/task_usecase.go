package usecases

import (
	"context"
	"fmt"
	"task_manager/domain"
	"time"
)

type TaskUseCase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUseCase(tr domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: tr,
	}
}

func (tuc *TaskUseCase) CreateTask(ctx context.Context, task domain.Task) error {
	if task.Title == "" {
		return fmt.Errorf("title required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := tuc.taskRepo.CreateTask(ctx, task); err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (tuc *TaskUseCase) UpdateTask(ctx context.Context, id string, task domain.Task) error {
	if id == "" {
		return fmt.Errorf("id required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := tuc.taskRepo.UpdateTask(ctx, id, task); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

func (tuc *TaskUseCase) DeleteTask(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("id required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := tuc.taskRepo.DeleteTask(ctx, id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}

func (tuc *TaskUseCase) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	alltask, err := tuc.taskRepo.GetAllTasks(ctx)
	if err != nil {
		return []domain.Task{}, fmt.Errorf("failed to fetch tasks: %w", err)
	}

	return alltask, err
}

func (tuc *TaskUseCase) GetTaskById(ctx context.Context, id string) (domain.Task, error) {
	if id == "" {
		return domain.Task{}, fmt.Errorf("id required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	task, err := tuc.taskRepo.GetTaskById(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to fetch task: %w", err)
	}

	return task, nil
}
