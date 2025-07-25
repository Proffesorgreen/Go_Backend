package data

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/zaahidali/task_manager_api/model"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTasks() ([]bson.M, error) {
	if Client == nil {
		return nil, fmt.Errorf("no connection established")
	}

	var tasks []bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskById(id int) (model.Task, error) {
	if Client == nil {
		return model.Task{}, fmt.Errorf("no connection established")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task model.Task
	filter := bson.M{"id": strconv.Itoa(id)}
	err := taskCollection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		return model.Task{}, err
	} else {
		return task, nil
	}
}

func UpdateTask(id int, updatedTask model.Task) (bson.M, error) {
	if Client == nil {
		return nil, fmt.Errorf("no connection established")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": strconv.Itoa(id)}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
		},
	}

	if res, err := taskCollection.UpdateOne(ctx, filter, update); err != nil {
		return nil, err
	} else {
		if res.MatchedCount == 0 {
			return nil, fmt.Errorf("no task found")
		}
		return update, nil
	}
}

func DeleteTask(id int) error {
	if Client == nil {
		return fmt.Errorf("no connection established")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": strconv.Itoa(id)}
	if _, err := taskCollection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

func CreateTask(newTask model.Task) error {
	if Client == nil {
		return fmt.Errorf("no connection established")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task model.Task
	filter := bson.M{"title": newTask.Title, "description": newTask.Description}
	if err := taskCollection.FindOne(ctx, filter).Decode(&task); err == nil {
		return fmt.Errorf("task already exists")
	}
	
	if _, err := taskCollection.InsertOne(ctx, newTask); err != nil {
		return err
	}
	return nil
}
