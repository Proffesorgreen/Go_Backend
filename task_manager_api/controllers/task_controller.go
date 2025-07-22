package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/data"
	"github.com/zaahidali/task_manager_api/model"
)

func HandleGetAllTasks(c *gin.Context) {
	Tasks := data.GetAllTasks()
	c.IndentedJSON(http.StatusOK, Tasks)
}

func HandleGetTaskById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Tasks, err := data.GetTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	} else {
		c.IndentedJSON(http.StatusOK, Tasks)
	}
}

func HandeleUpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTask model.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	newTask, err := data.UpdateTask(id, updatedTask)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
	} else {
		c.IndentedJSON(http.StatusOK, newTask)
	}
}

func HandleDeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := data.DeleteTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not Found"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully deleted task"})
	}
}

func HandleCreateTask(c *gin.Context) {
	var newtask model.Task
	if err := c.BindJSON(&newtask); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	data.CreateTask(newtask)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Added"})
}
