package controller

import (
	"errors"
	"net/http"
	"task_manager/domain"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

type TaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

type TaskDeleteRequest struct {
	ID string `json:"id"`
}

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AllController struct {
	userusecase domain.UserUseCase
	taskusecase domain.TaskUseCase
}

func NewAllController(uuc domain.UserUseCase, tuc domain.TaskUseCase) *AllController {
	return &AllController{
		userusecase: uuc,
		taskusecase: tuc,
	}
}

func ToDTO(u *domain.Task) *TaskResponse {
	return &TaskResponse{
		ID:          u.ID,
		Title:       u.Title,
		Description: u.Description,
		Status:      u.Status,
		DueDate:     u.DueDate,
	}
}

func ToDomain(u *TaskRequest) *domain.Task {
	return &domain.Task{
		Title:       u.Title,
		Description: u.Description,
		Status:      u.Status,
		DueDate:     u.DueDate,
	}
}

func (ac *AllController) HandleRegister(c *gin.Context) {
	var req UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect request"})
		return
	}

	if err := ac.userusecase.RegisterUser(c.Request.Context(), req.Email, req.Username, req.Password); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered sucessfully"})
}

func (ac *AllController) HandleLogin(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect request"})
		return
	}

	accessToken, err := ac.userusecase.LoginUser(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func (ac *AllController) HandleCreateTask(c *gin.Context) {
	var req TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect request"})
		return
	}

	if err := ac.taskusecase.CreateTask(c.Request.Context(), *ToDomain(&req)); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "task created sucessfully"})
}

func (ac *AllController) HandleUpdateTask(c *gin.Context) {
	id := c.Param("id")
	var req TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect request"})
		return
	}

	if err := ac.taskusecase.UpdateTask(c.Request.Context(), id, *ToDomain(&req)); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task updated successfully",
	})
}

func (ac *AllController) HandleDeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := ac.taskusecase.DeleteTask(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task deleted sucessfully"})
}

func (ac *AllController) HandleGetAllTasks(c *gin.Context) {
	res, err := ac.taskusecase.GetAllTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	taskres := make([]TaskResponse, 0, len(res))
	for i := range res {
		taskres = append(taskres, *ToDTO(&res[i]))
	}

	c.JSON(http.StatusOK, gin.H{"tasks": taskres})
}

func (ac *AllController) HandleGetTaskById(c *gin.Context) {
	id := c.Param("id")
	res, err := ac.taskusecase.GetTaskById(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Task not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": *ToDTO(&res)})
}
