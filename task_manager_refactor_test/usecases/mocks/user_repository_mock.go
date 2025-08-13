package mocks

import (
	"context"
	"task_manager/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) AddUser(ctx context.Context, user domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, Email string) (domain.User, error) {
	args := m.Called(ctx, Email)
	if args.Get(0) == nil {
		return domain.User{}, args.Error(1)
	}
	return args.Get(0).(domain.User), args.Error(1)
}
