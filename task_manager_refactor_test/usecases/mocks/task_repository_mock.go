package mocks

import (
	"context"
	"task_manager/domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(ctx context.Context, newTask domain.Task) error {
	args := m.Called(ctx, newTask)
	return args.Error(0)
}

func (m *MockTaskRepository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return []domain.Task{}, args.Error(1)
	}
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTaskById(ctx context.Context, id string) (domain.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return domain.Task{}, args.Error(1)
	}
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(ctx context.Context, id string, updatedTask domain.Task) error {
	args := m.Called(ctx, id, updatedTask)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
