package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) RegisterUser(ctx context.Context, email, username, password string) error {
	args := m.Called(ctx, email, username, password)
	return args.Error(0)
}

func (m *MockUserUseCase) LoginUser(ctx context.Context, email, password string) (string, error) {
	args := m.Called(ctx, email, password)
	return args.String(0), args.Error(1)
}
