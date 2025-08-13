package infrastructure

import (
	"task_manager/domain"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	secret := "supersecretkey"
	jp := NewJwtProvider(secret)

	user := domain.User{
		ID:       "user123",
		Username: "testuser",
		Role:     domain.RoleUser,
	}

	token, err := jp.GenerateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate the generated token
	claims, err := jp.ValidateJWT(token)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, claims.UserId)
	assert.Equal(t, user.Username, claims.Username)
	assert.Equal(t, user.Role, claims.Role)
	assert.True(t, claims.ExpiresAt > time.Now().Unix())
}

func TestValidateJWT(t *testing.T) {
	secret := "supersecretkey"
	jp := NewJwtProvider(secret)

	user := domain.User{
		ID:       "user123",
		Username: "testuser",
		Role:     domain.RoleUser,
	}

	validToken, err := jp.GenerateToken(user)
	assert.NoError(t, err)

	t.Run("valid token", func(t *testing.T) {
		claims, err := jp.ValidateJWT(validToken)
		assert.NoError(t, err)
		assert.Equal(t, user.ID, claims.UserId)
		assert.Equal(t, user.Username, claims.Username)
		assert.Equal(t, user.Role, claims.Role)
	})

	t.Run("invalid token", func(t *testing.T) {
		invalidToken := "invalid.token.string"
		_, err := jp.ValidateJWT(invalidToken)
		assert.Error(t, err)
	})

	t.Run("token with wrong secret", func(t *testing.T) {
		wrongSecretProvider := NewJwtProvider("wrongsecret")
		wrongSecretToken, err := wrongSecretProvider.GenerateToken(user)
		assert.NoError(t, err)

		_, err = jp.ValidateJWT(wrongSecretToken)
		assert.Error(t, err)
	})

	t.Run("expired token", func(t *testing.T) {
		expiredClaims := domain.Claims{
			UserId:   user.ID,
			Username: user.Username,
			Role:     user.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(),
				IssuedAt:  time.Now().Add(-2 * time.Hour).Unix(),
			},
		}

		expiredToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims).SignedString([]byte(secret))
		assert.NoError(t, err)

		_, err = jp.ValidateJWT(expiredToken)
		assert.Error(t, err)
	})
}
