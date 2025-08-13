package mocks

import (
	"task_manager/domain"

	"github.com/stretchr/testify/mock"
)

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateToken(user domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockJWTService) ValidateJWT(claims string) (domain.Claims, error) {
	args := m.Called(claims)
	if args.Get(0) == nil {
		return domain.Claims{}, args.Error(1)
	}
	return args.Get(0).(domain.Claims), args.Error(1)
}
