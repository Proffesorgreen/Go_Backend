package data

import (
	"fmt"
	"strconv"
	"time"

	"github.com/zaahidali/task_manager_api/model"
)

var Tasks = []model.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func GetAllTasks() []model.Task {
	return Tasks
}

func GetTaskById(id int) (model.Task, error) {
	for _, task := range Tasks {
		if val, _ := strconv.Atoi(task.ID); val == id {
			return task, nil
		}
	}
	return model.Task{}, fmt.Errorf("something went wrong")
}

func UpdateTask(id int, updatedTask model.Task) (model.Task, error) {
	for idx := range Tasks {
		if val, _ := strconv.Atoi(Tasks[idx].ID); val == id {
			if updatedTask.Title != "" {
				Tasks[idx].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				Tasks[idx].Description = updatedTask.Description
			}
			updatedTask = Tasks[idx]
			return updatedTask, nil
		}
	}
	return model.Task{}, fmt.Errorf("something went wrong")
}

func DeleteTask(id int) error {
	for idx := range Tasks {
		if val, _ := strconv.Atoi(Tasks[idx].ID); val == id {
			Tasks = append(Tasks[:idx], Tasks[idx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("something went wrong")
}

func CreateTask(newTask model.Task) {
	Tasks = append(Tasks, newTask)
}
