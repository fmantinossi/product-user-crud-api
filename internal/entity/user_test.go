package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Contains(t, user.Email, "@")
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe Jr", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.Nil(t, user.ValidatePassword("123456"))
	assert.NotNil(t, user.ValidatePassword("aaaaaaa"))
	assert.NotEqual(t, "123456", user.Password)
}
