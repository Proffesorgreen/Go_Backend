package mocks

import (
	"context"
	"task_manager/domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskUseCase struct {
	mock.Mock
}

func (m *MockTaskUseCase) CreateTask(ctx context.Context, task domain.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *MockTaskUseCase) UpdateTask(ctx context.Context, id string, task domain.Task) error {
	args := m.Called(ctx, id, task)
	return args.Error(0)
}

func (m *MockTaskUseCase) DeleteTask(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTaskUseCase) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) GetTaskById(ctx context.Context, id string) (domain.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return domain.Task{}, args.Error(1)
	}
	return args.Get(0).(domain.Task), args.Error(1)
}
