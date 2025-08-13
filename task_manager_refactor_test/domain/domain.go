package domain

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

type Role string

var (
	RoleAdmin Role = "Admin"
	RoleUser  Role = "User"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Status      string
	DueDate     time.Time
}

type User struct {
	ID       string
	Email    string
	Username string
	Password string
	Role     Role
}

type Claims struct {
	UserId   string
	Role     Role
	Username string
	jwt.StandardClaims
}

type UserUseCase interface {
	RegisterUser(ctx context.Context, email, username, password string) error
	LoginUser(ctx context.Context, email, password string) (string, error)
}

type TaskUseCase interface {
	CreateTask(ctx context.Context, task Task) error
	UpdateTask(ctx context.Context, id string, task Task) error
	DeleteTask(ctx context.Context, id string) error
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetTaskById(ctx context.Context, id string) (Task, error)
}

type UserRepository interface {
	AddUser(ctx context.Context, user User) error
	GetUserByEmail(ctx context.Context, Email string) (User, error)
}

type TaskRepository interface {
	CreateTask(ctx context.Context, newTask Task) error
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetTaskById(ctx context.Context, id string) (Task, error)
	UpdateTask(ctx context.Context, id string, updatedTask Task) error
	DeleteTask(ctx context.Context, id string) error
}

type JWT_Service interface {
	GenerateToken(user User) (string, error)
	ValidateJWT(claims string) (Claims, error)
}

type Passwordservice interface {
	HashingPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}
