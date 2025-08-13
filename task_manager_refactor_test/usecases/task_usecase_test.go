package usecases

import (
	"context"
	"errors"
	"task_manager/domain"
	"task_manager/usecases/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTask(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	u := NewTaskUseCase(mockTaskRepo)

	t.Run("success", func(t *testing.T) {
		task := domain.Task{
			Title:       "Test Task",
			Description: "Test Description",
			Status:      "Pending",
			DueDate:     time.Now(),
		}
		mockTaskRepo.On("CreateTask", mock.Anything, task).Return(nil).Once()

		err := u.CreateTask(context.Background(), task)

		assert.NoError(t, err)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("empty title", func(t *testing.T) {
		task := domain.Task{
			Title: "",
		}

		err := u.CreateTask(context.Background(), task)

		assert.Error(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		task := domain.Task{
			Title: "Test Task",
		}
		mockTaskRepo.On("CreateTask", mock.Anything, task).Return(errors.New("repository error")).Once()

		err := u.CreateTask(context.Background(), task)

		assert.Error(t, err)
		mockTaskRepo.AssertExpectations(t)
	})
}

func TestGetTaskById(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	u := NewTaskUseCase(mockTaskRepo)

	t.Run("success", func(t *testing.T) {
		task := domain.Task{
			ID:    "1",
			Title: "Test Task",
		}
		mockTaskRepo.On("GetTaskById", mock.Anything, "1").Return(task, nil).Once()

		result, err := u.GetTaskById(context.Background(), "1")

		assert.NoError(t, err)
		assert.Equal(t, task, result)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockTaskRepo.On("GetTaskById", mock.Anything, "1").Return(domain.Task{}, errors.New("not found")).Once()

		_, err := u.GetTaskById(context.Background(), "1")

		assert.Error(t, err)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("empty id", func(t *testing.T) {
		_, err := u.GetTaskById(context.Background(), "")

		assert.Error(t, err)
	})
}

func TestUpdateTask(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	u := NewTaskUseCase(mockTaskRepo)

	t.Run("success", func(t *testing.T) {
		id := "some-id"
		task := domain.Task{
			Title:       "Updated Task",
			Description: "Updated Description",
			Status:      "Completed",
			DueDate:     time.Now(),
		}
		mockTaskRepo.On("UpdateTask", mock.Anything, id, task).Return(nil).Once()

		err := u.UpdateTask(context.Background(), id, task)

		assert.NoError(t, err)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("empty id", func(t *testing.T) {
		task := domain.Task{
			Title: "Updated Task",
		}
		err := u.UpdateTask(context.Background(), "", task)

		assert.Error(t, err)
		assert.EqualError(t, err, "id required")
	})

	t.Run("repository error", func(t *testing.T) {
		id := "some-id"
		task := domain.Task{
			Title: "Updated Task",
		}
		mockTaskRepo.On("UpdateTask", mock.Anything, id, task).Return(errors.New("repository error")).Once()

		err := u.UpdateTask(context.Background(), id, task)

		assert.Error(t, err)
		mockTaskRepo.AssertExpectations(t)
	})
}

func TestDeleteTask(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	u := NewTaskUseCase(mockTaskRepo)

	t.Run("success", func(t *testing.T) {
		id := "some-id"
		mockTaskRepo.On("DeleteTask", mock.Anything, id).Return(nil).Once()

		err := u.DeleteTask(context.Background(), id)

		assert.NoError(t, err)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("empty id", func(t *testing.T) {
		err := u.DeleteTask(context.Background(), "")

		assert.Error(t, err)
		assert.EqualError(t, err, "id required")
	})

	t.Run("repository error", func(t *testing.T) {
		id := "some-id"
		mockTaskRepo.On("DeleteTask", mock.Anything, id).Return(errors.New("repository error")).Once()

		err := u.DeleteTask(context.Background(), id)

		assert.Error(t, err)
		mockTaskRepo.AssertExpectations(t)
	})
}

func TestGetAllTasks(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	u := NewTaskUseCase(mockTaskRepo)

	t.Run("success", func(t *testing.T) {
		tasks := []domain.Task{
			{ID: "1", Title: "Task 1"},
			{ID: "2", Title: "Task 2"},
		}
		mockTaskRepo.On("GetAllTasks", mock.Anything).Return(tasks, nil).Once()

		result, err := u.GetAllTasks(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, tasks, result)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		mockTaskRepo.On("GetAllTasks", mock.Anything).Return(nil, errors.New("repository error")).Once()

		_, err := u.GetAllTasks(context.Background())

		assert.Error(t, err)
		mockTaskRepo.AssertExpectations(t)
	})
}

// func TestGetTaskById(t *testing.T) {
// 	mockTaskRepo := new(mocks.MockTaskRepository)
// 	u := NewTaskUseCase(mockTaskRepo)

// 	t.Run("success", func(t *testing.T) {
// 		task := domain.Task{
// 			ID:    "1",
// 			Title: "Test Task",
// 		}
// 		mockTaskRepo.On("GetTaskById", mock.Anything, "1").Return(task, nil).Once()

// 		result, err := u.GetTaskById(context.Background(), "1")

// 		assert.NoError(t, err)
// 		assert.Equal(t, task, result)
// 		mockTaskRepo.AssertExpectations(t)
// 	})

// 	t.Run("not found", func(t *testing.T) {
// 		mockTaskRepo.On("GetTaskById", mock.Anything, "1").Return(domain.Task{}, errors.New("not found")).Once()

// 		_, err := u.GetTaskById(context.Background(), "1")

// 		assert.Error(t, err)
// 		mockTaskRepo.AssertExpectations(t)
// 	})

// 	t.Run("empty id", func(t *testing.T) {
// 		_, err := u.GetTaskById(context.Background(), "")

// 		assert.Error(t, err)
// 	})
// }
