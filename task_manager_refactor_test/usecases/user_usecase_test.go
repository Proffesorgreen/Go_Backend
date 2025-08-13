package usecases

import (
    "context"
    "errors"
    "testing"
    "task_manager/domain"
    "task_manager/usecases/mocks"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
    // Create instances of our mocks
    mockUserRepo := new(mocks.MockUserRepository)
    mockPasswordService := new(mocks.MockPasswordService)
    mockJWTService := new(mocks.MockJWTService)

    // Create an instance of our use case
    u := NewUserUseCase(mockUserRepo, mockPasswordService, mockJWTService)

    t.Run("success", func(t *testing.T) {
        email := "test@example.com"
        username := "testuser"
        password := "password"
        hashedPassword := "hashed_password"

        // Setup the expectations
        mockUserRepo.On("GetUserByEmail", mock.Anything, email).
            Return(domain.User{}, errors.New("user not found")).Once()
        mockPasswordService.On("HashingPassword", password).
            Return(hashedPassword, nil).Once()
        mockUserRepo.On("AddUser", mock.Anything, mock.AnythingOfType("domain.User")).
            Return(nil).Once()

        // Call the method we want to test
        err := u.RegisterUser(context.Background(), email, username, password)

        // Assert that the expectations were met
        mockUserRepo.AssertExpectations(t)
        mockPasswordService.AssertExpectations(t)
        assert.NoError(t, err)
    })

    t.Run("user already exists", func(t *testing.T) {
        email := "test@example.com"
        username := "testuser"
        password := "password"

        // Setup the expectations
        mockUserRepo.On("GetUserByEmail", mock.Anything, email).
            Return(domain.User{}, nil).Once()

        // Call the method we want to test
        err := u.RegisterUser(context.Background(), email, username, password)

        // Assert that the expectations were met
        mockUserRepo.AssertExpectations(t)
        assert.Error(t, err)
    })

    t.Run("password hashing error", func(t *testing.T) {
        email := "test@example.com"
        username := "testuser"
        password := "password"

        // Setup the expectations
        mockUserRepo.On("GetUserByEmail", mock.Anything, email).
            Return(domain.User{}, errors.New("user not found")).Once()
        mockPasswordService.On("HashingPassword", password).
            Return("", errors.New("hashing error")).Once()

        // Call the method we want to test
        err := u.RegisterUser(context.Background(), email, username, password)

        // Assert that the expectations were met
        mockUserRepo.AssertExpectations(t)
        mockPasswordService.AssertExpectations(t)
        assert.Error(t, err)
    })

    t.Run("add user error", func(t *testing.T) {
        email := "test@example.com"
        username := "testuser"
        password := "password"
        hashedPassword := "hashed_password"

        // Setup the expectations
        mockUserRepo.On("GetUserByEmail", mock.Anything, email).
            Return(domain.User{}, errors.New("user not found")).Once()
        mockPasswordService.On("HashingPassword", password).
            Return(hashedPassword, nil).Once()
        mockUserRepo.On("AddUser", mock.Anything, mock.AnythingOfType("domain.User")).
            Return(errors.New("add user error")).Once()

        // Call the method we want to test
        err := u.RegisterUser(context.Background(), email, username, password)

        // Assert that the expectations were met
        mockUserRepo.AssertExpectations(t)
        mockPasswordService.AssertExpectations(t)
        assert.Error(t, err)
    })
}
