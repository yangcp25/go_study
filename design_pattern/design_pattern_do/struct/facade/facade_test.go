package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_UserLoginOrRegister(t *testing.T) {
	service := UserService{}
	user, err := service.UserLoginOrRegister("13001010101", "1234")
	assert.NoError(t, err)
	assert.Equal(t, &User{name: "13001010101"}, user)
}
