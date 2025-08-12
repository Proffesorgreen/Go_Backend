package usecases

import (
	"context"
	"task_manager/domain"
)

type UserUseCase struct {
	userRepo domain.UserRepository
	userPass domain.Passwordservice
	userJWT  domain.JWT_Service
}

func NewUserUseCase(ur domain.UserRepository, up domain.Passwordservice, uj domain.JWT_Service) *UserUseCase {
	return &UserUseCase{
		userRepo: ur,
		userPass: up,
		userJWT:  uj,
	}
}

func (uuc *UserUseCase) RegisterUser(ctx context.Context, email, username, password string) error {
	if _, err := uuc.userRepo.GetUserByEmail(ctx, username); err == nil {
		return err
	}

	hashedPassword, err := uuc.userPass.HashingPassword(password)
	if err != nil {
		return err
	}

	user := domain.User{
		Email:    email,
		Username: username,
		Password: hashedPassword,
		Role:     domain.RoleUser,
	}

	if err := uuc.userRepo.AddUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (uuc *UserUseCase) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := uuc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := uuc.userPass.ComparePassword(user.Password, password); err != nil {
		return "", err
	}

	token, er := uuc.userJWT.GenerateToken(user)
	if er != nil {
		return "", err
	}

	return token, nil
}
