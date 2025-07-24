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
	if tasks, err := data.GetTaskById(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	} else {
		c.IndentedJSON(http.StatusOK, tasks)
	}
}

func HandeleUpdateTask(c *gin.Context) {
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
	err := data.CreateTask(newtask)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Added"})
}
