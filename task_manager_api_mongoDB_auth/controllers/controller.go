package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/data"
	"github.com/zaahidali/task_manager_api/model"
)

func HandleGetAllTasks(c *gin.Context) {
	if tasks, err := data.GetAllTasks(); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	} else {
		c.IndentedJSON(http.StatusOK, tasks)
	}
}

func HandleGetTaskById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if task, err := data.GetTaskById(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	} else {
		c.IndentedJSON(http.StatusOK, task)
	}
}

func HandleUpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTask model.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	if newTask, err := data.UpdateTask(id, updatedTask); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	} else {
		c.IndentedJSON(http.StatusOK, newTask)
	}
}

func HandleDeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := data.DeleteTask(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully deleted task"})
}

func HandleCreateTask(c *gin.Context) {
	var newtask model.Task
	if err := c.BindJSON(&newtask); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := data.CreateTask(newtask); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Added"})
}

func HandleRegister(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := data.RegisterUser(user); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User Added"})
}

func HandleLogin(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	token, err := data.LoginUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}
