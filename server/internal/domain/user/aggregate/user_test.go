package aggregate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	registerModel := UserRegisterModel{
		Username: "user",
		Email:    "email@domain.com",
		Password: "password",
	}

	user := NewUser(registerModel)

	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, registerModel.Username, user.Username)
	assert.Equal(t, registerModel.Email, user.Email)
	assert.NotEqual(t, registerModel.Password, user.Password)
	assert.NotZero(t, user.CreatedAt)
	assert.NotZero(t, user.UpdatedAt)
}

func TestUpdateEmail(t *testing.T) {
	user := &User{
		Email: "old@domain.com",
	}

	newEmail := "new@domain.com"
	user.UpdateEmail(newEmail)

	assert.Equal(t, newEmail, user.Email)
	assert.NotZero(t, user.UpdatedAt)
}
