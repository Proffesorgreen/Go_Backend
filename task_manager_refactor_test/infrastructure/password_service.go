package infrastructure

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordProvider struct {
	cost int
}

func NewPasswordProvider(cst int) *PasswordProvider {
	if cst == 0 {
		cst = bcrypt.DefaultCost
	}
	return &PasswordProvider{
		cost: cst,
	}
}

func (pp *PasswordProvider) HashingPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), pp.cost)
	if err != nil {
		return "", fmt.Errorf("password service: failed to generate hash: %w", err)
	}
	return string(hashedPassword), nil
}

func (pp *PasswordProvider) ComparePassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return fmt.Errorf("password service: password comparison failed: %w", err)
	}
	return nil
}
