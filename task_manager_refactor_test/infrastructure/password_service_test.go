package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashingPassword(t *testing.T) {
	pp := NewPasswordProvider(10) // Use default cost

	password := "mysecretpassword"
	hashedPassword, err := pp.HashingPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	// Verify that the hashed password can be compared successfully
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	assert.NoError(t, err)
}

func TestComparePassword(t *testing.T) {
	pp := NewPasswordProvider(10) // Use default cost

	password := "mysecretpassword"
	hashedPassword, _ := pp.HashingPassword(password)

	t.Run("correct password", func(t *testing.T) {
		err := pp.ComparePassword(hashedPassword, password)
		assert.NoError(t, err)
	})

	t.Run("incorrect password", func(t *testing.T) {
		err := pp.ComparePassword(hashedPassword, "wrongpassword")
		assert.Error(t, err)
		assert.EqualError(t, err, "password service: password comparison failed: crypto/bcrypt: hashedPassword is not the hash of the given password")
	})

	t.Run("empty hashed password", func(t *testing.T) {
		err := pp.ComparePassword("", password)
		assert.Error(t, err)
	})

	t.Run("empty plain password", func(t *testing.T) {
		err := pp.ComparePassword(hashedPassword, "")
		assert.Error(t, err)
	})
}
