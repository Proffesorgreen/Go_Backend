package infrastructure

import (
	"fmt"
	"task_manager/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwt_provider struct {
	SecretKey string
}

func NewJwtProvider(secret string) *jwt_provider {
	return &jwt_provider{
		SecretKey: secret,
	}
}

func (jp *jwt_provider) GenerateToken(user domain.User) (string, error) {
	claims := domain.Claims{
		UserId:   user.ID,
		Role:     user.Role,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jp.SecretKey))
}

func (jp *jwt_provider) ValidateJWT(tokenString string) (domain.Claims, error) {
	claims := &domain.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jp.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return domain.Claims{}, err
	}

	return *claims, nil
}
