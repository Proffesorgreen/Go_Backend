package data

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zaahidali/task_manager_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user model.User) error {
	if Client == nil {
		return fmt.Errorf("no connection established")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var existinguser model.User
	filter := bson.M{"username": user.Username}
	if err := userCollection.FindOne(ctx, filter).Decode(&existinguser); err == nil {
		return fmt.Errorf("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if _, err := userCollection.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}

func LoginUser(user model.User) (string, error) {
	if Client == nil {
		return "", fmt.Errorf("no connection established")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var existinguser model.User
	filter := bson.M{"username": user.Username}
	if err := userCollection.FindOne(ctx, filter).Decode(&existinguser); err != nil {
		return "", fmt.Errorf("unregistered user")
	}

	if bcrypt.CompareHashAndPassword([]byte(existinguser.Password), []byte(user.Password)) != nil {
		return "", fmt.Errorf("incorrect password")
	}

	jwt_secret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role":  existinguser.Role,
		"email": existinguser.Email,
	})

	signedToken, ok := token.SignedString(jwt_secret)
	if ok != nil {
		return "", fmt.Errorf("internal error")
	}

	return signedToken, nil
}
