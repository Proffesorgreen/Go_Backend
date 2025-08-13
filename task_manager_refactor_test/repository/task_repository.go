package repository

import (
	"context"
	"task_manager/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepoDTO struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	DueDate     time.Time          `bson:"due_date"`
}

func ToDomainTask(u *TaskRepoDTO) *domain.Task {
	return &domain.Task{
		ID:          u.ID.Hex(),
		Title:       u.Title,
		Description: u.Description,
		Status:      u.Status,
		DueDate:     u.DueDate,
	}
}

func ToDTOTask(u *domain.Task) *TaskRepoDTO {
	id := primitive.NewObjectID()
	return &TaskRepoDTO{
		ID:          id,
		Title:       u.Title,
		Description: u.Description,
		Status:      u.Status,
		DueDate:     u.DueDate,
	}
}

type TaskRepoService struct {
	collection *mongo.Collection
}

func NewTaskRepoService(coll *mongo.Collection) *TaskRepoService {
	return &TaskRepoService{
		collection: coll,
	}
}

func (trs *TaskRepoService) CreateTask(ctx context.Context, newTask domain.Task) error {
	if _, err := trs.collection.InsertOne(ctx, *ToDTOTask(&newTask)); err != nil {
		return err
	}
	return nil
}

func (trs *TaskRepoService) UpdateTask(ctx context.Context, id string, updatedTask domain.Task) error {
	newid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": newid}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"status":      updatedTask.Status,
			"due_date":    updatedTask.DueDate,
		},
	}
	if _, err := trs.collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func (trs *TaskRepoService) DeleteTask(ctx context.Context, id string) error {
	newid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": newid}
	if _, err := trs.collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (trs *TaskRepoService) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	dtotask := &[]TaskRepoDTO{}

	cursor, err := trs.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, dtotask); err != nil {
		return nil, err
	}

	domaintask := make([]domain.Task, len(*dtotask))
	for i, dto := range *dtotask {
		domaintask[i] = *ToDomainTask(&dto)
	}

	return domaintask, nil
}

func (trs *TaskRepoService) GetTaskById(ctx context.Context, id string) (domain.Task, error) {
	var res TaskRepoDTO
	newid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}

	filter := bson.M{"_id": newid}
	if err := trs.collection.FindOne(ctx, filter).Decode(&res); err != nil {
		return domain.Task{}, err
	}

	return *ToDomainTask(&res), nil
}
