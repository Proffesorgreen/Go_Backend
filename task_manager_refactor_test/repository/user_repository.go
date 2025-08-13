package repository

import (
	"context"
	"errors"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoDTO struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}

func ToDTO(u *domain.User) *UserRepoDTO {
	id := primitive.NewObjectID()
	return &UserRepoDTO{
		ID:       id,
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Role:     string(u.Role),
	}
}

func ToDomain(u *UserRepoDTO) *domain.User {
	return &domain.User{
		ID:       u.ID.Hex(),
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Role:     domain.Role(u.Role),
	}
}

type UserRepoService struct {
	collection *mongo.Collection
}

func NewUserRepoService(coll *mongo.Collection) *UserRepoService {
	return &UserRepoService{
		collection: coll,
	}
}

func (ups *UserRepoService) AddUser(ctx context.Context, user domain.User) error {
	if _, err := ups.collection.InsertOne(ctx, *ToDTO(&user)); err != nil {
		return err
	}
	return nil
}

func (ups *UserRepoService) GetUserByEmail(ctx context.Context, Email string) (domain.User, error) {
	var res UserRepoDTO
	filter := bson.M{"email": Email}
	if err := ups.collection.FindOne(ctx, filter).Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return *ToDomain(&res), nil
}
