package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	usecasemocks "task_manager/delivery/controller/mocks"
	"task_manager/domain"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestHandleRegister(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase) // Required for controller construction
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		registerReq := UserRegisterRequest{
			Email:    "test@example.com",
			Username: "testuser",
			Password: "password",
		}
		body, _ := json.Marshal(registerReq)

		mockUserUseCase.On("RegisterUser", mock.Anything, registerReq.Email, registerReq.Username, registerReq.Password).
			Return(nil).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleRegister(c)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "user registered successfully")
		mockUserUseCase.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		fmt.Println("invalid request body")
		c.Request = httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(`{"email": "test2@example.com", "username":`))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleRegister(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "incorrect request")
		mockUserUseCase.AssertNotCalled(t, "RegisterUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run("usecase returns error", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase) // Required for controller construction
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		registerReq := UserRegisterRequest{
			Email:    "test@example.com",
			Username: "testuser",
			Password: "password",
		}
		body, _ := json.Marshal(registerReq)

		mockUserUseCase.On("RegisterUser", mock.Anything, registerReq.Email, registerReq.Username, registerReq.Password).
			Return(errors.New("user already exists")).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleRegister(c)

		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Contains(t, w.Body.String(), "user already exists")
		mockUserUseCase.AssertExpectations(t)
	})
}

func TestHandleLogin(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		loginReq := UserLoginRequest{
			Email:    "test@example.com",
			Password: "password",
		}
		body, _ := json.Marshal(loginReq)
		accessToken := "mock_access_token"

		mockUserUseCase.On("LoginUser", mock.Anything, loginReq.Email, loginReq.Password).
			Return(accessToken, nil).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleLogin(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), accessToken)
		mockUserUseCase.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase) // Required for controller construction
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewBufferString("invalid json"))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleLogin(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "incorrect request")
		mockUserUseCase.AssertNotCalled(t, "LoginUser", mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run("usecase returns error", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		loginReq := UserLoginRequest{
			Email:    "test@example.com",
			Password: "password",
		}
		body, _ := json.Marshal(loginReq)

		mockUserUseCase.On("LoginUser", mock.Anything, loginReq.Email, loginReq.Password).
			Return("", errors.New("invalid credentials")).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleLogin(c)

		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Contains(t, w.Body.String(), "invalid credentials")
		mockUserUseCase.AssertExpectations(t)
	})
}

func TestHandleDeleteTask(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		taskID := "123"
		mockTaskUseCase.On("DeleteTask", mock.Anything, taskID).Return(nil).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: taskID}}
		c.Request = httptest.NewRequest(http.MethodDelete, "/tasks/"+taskID, nil)

		controller.HandleDeleteTask(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "task deleted successfully")
		mockTaskUseCase.AssertExpectations(t)
	})

	t.Run("usecase returns error", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		taskID := "123"
		mockTaskUseCase.On("DeleteTask", mock.Anything, taskID).Return(errors.New("task not found")).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: taskID}}
		c.Request = httptest.NewRequest(http.MethodDelete, "/tasks/"+taskID, nil)

		controller.HandleDeleteTask(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockTaskUseCase.AssertExpectations(t)
	})
}

func TestHandleGetAllTasks(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		tasks := []domain.Task{
			{ID: "1", Title: "Task 1", Description: "Desc 1", Status: "Pending", DueDate: time.Now()},
			{ID: "2", Title: "Task 2", Description: "Desc 2", Status: "Completed", DueDate: time.Now()},
		}
		mockTaskUseCase.On("GetAllTasks", mock.Anything).Return(tasks, nil).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/tasks", nil)

		controller.HandleGetAllTasks(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string][]TaskResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response["tasks"], 2)
		assert.Equal(t, tasks[0].Title, response["tasks"][0].Title)
		mockTaskUseCase.AssertExpectations(t)
	})

	t.Run("usecase returns error", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		mockTaskUseCase.On("GetAllTasks", mock.Anything).Return(nil, errors.New("database error")).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/tasks", nil)

		controller.HandleGetAllTasks(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockTaskUseCase.AssertExpectations(t)
	})
}

func TestHandleGetTaskById(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		taskID := "123"
		task := domain.Task{ID: taskID, Title: "Found Task", Description: "Desc", Status: "Pending", DueDate: time.Now()}
		mockTaskUseCase.On("GetTaskById", mock.Anything, taskID).Return(task, nil).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: taskID}}
		c.Request = httptest.NewRequest(http.MethodGet, "/tasks/"+taskID, nil)

		controller.HandleGetTaskById(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]TaskResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, task.Title, response["task"].Title)
		mockTaskUseCase.AssertExpectations(t)
	})

	t.Run("task not found", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		taskID := "123"
		mockTaskUseCase.On("GetTaskById", mock.Anything, taskID).Return(domain.Task{}, mongo.ErrNoDocuments).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: taskID}}
		c.Request = httptest.NewRequest(http.MethodGet, "/tasks/"+taskID, nil)

		controller.HandleGetTaskById(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Task not found")
		mockTaskUseCase.AssertExpectations(t)
	})

	t.Run("usecase returns generic error", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		taskID := "123"
		mockTaskUseCase.On("GetTaskById", mock.Anything, taskID).Return(domain.Task{}, errors.New("some other error")).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: taskID}}
		c.Request = httptest.NewRequest(http.MethodGet, "/tasks/"+taskID, nil)

		controller.HandleGetTaskById(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "internal server error")
		mockTaskUseCase.AssertExpectations(t)
	})
}

func TestHandleCreateTask(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		taskReq := TaskRequest{
			Title:       "New Task",
			Description: "Description",
			Status:      "Pending",
			DueDate:     time.Now(),
		}
		body, _ := json.Marshal(taskReq)

		mockTaskUseCase.On("CreateTask", mock.Anything, mock.AnythingOfType("domain.Task")).Return(nil).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleCreateTask(c)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "task created successfully")
		mockTaskUseCase.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString("invalid json"))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleCreateTask(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "incorrect request")
		mockTaskUseCase.AssertNotCalled(t, "CreateTask", mock.Anything, mock.Anything)
	})

	t.Run("usecase returns error", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		taskReq := TaskRequest{
			Title: "New Task",
		}
		body, _ := json.Marshal(taskReq)

		mockTaskUseCase.On("CreateTask", mock.Anything, mock.AnythingOfType("domain.Task")).
			Return(errors.New("title required")).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleCreateTask(c)

		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Contains(t, w.Body.String(), "title required")
		mockTaskUseCase.AssertExpectations(t)
	})
}

func TestHandleUpdateTask(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		taskID := "123"
		taskReq := TaskRequest{
			Title:       "Updated Task",
			Description: "Updated Description",
			Status:      "Completed",
			DueDate:     time.Now(),
		}
		body, _ := json.Marshal(taskReq)

		mockTaskUseCase.On("UpdateTask", mock.Anything, taskID, mock.AnythingOfType("domain.Task")).Return(nil).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: taskID}}
		c.Request = httptest.NewRequest(http.MethodPatch, "/tasks/"+taskID, bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleUpdateTask(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "task updated successfully")
		mockTaskUseCase.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		taskID := "123"
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: taskID}}
		c.Request = httptest.NewRequest(http.MethodPatch, "/tasks/"+taskID, bytes.NewBufferString("invalid json"))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleUpdateTask(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "incorrect request")
		mockTaskUseCase.AssertNotCalled(t, "UpdateTask", mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run("usecase returns error", func(t *testing.T) {
		mockUserUseCase := new(usecasemocks.MockUserUseCase)
		mockTaskUseCase := new(usecasemocks.MockTaskUseCase)
		controller := NewAllController(mockUserUseCase, mockTaskUseCase)

		taskID := "123"
		taskReq := TaskRequest{
			Title: "Updated Task",
		}
		body, _ := json.Marshal(taskReq)

		mockTaskUseCase.On("UpdateTask", mock.Anything, taskID, mock.AnythingOfType("domain.Task")).
			Return(errors.New("task not found")).Once()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: taskID}}
		c.Request = httptest.NewRequest(http.MethodPatch, "/tasks/"+taskID, bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.HandleUpdateTask(c)

		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Contains(t, w.Body.String(), "task not found")
		mockTaskUseCase.AssertExpectations(t)
	})
}
