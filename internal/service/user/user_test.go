package user

import (
	"fmt"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/felixlambertv/go-cleanplate/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userRepoMock = new(mocks.IUserRepo)
var userService = NewUserService(userRepoMock)
var userDummy = []model.User{
	{
		ID:       1,
		Name:     "",
		Email:    "",
		Password: "",
	},
	{
		ID:       1,
		Name:     "",
		Email:    "",
		Password: "",
	},
}

func TestMain(m *testing.M) {
	fmt.Print("before")

	m.Run()
	fmt.Println("after")
}

func TestUser_GetUsers(t *testing.T) {
	userRepoMock.On("FindAll").Return(userDummy, nil)
	users, err := userService.GetUsers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
	assert.Equal(t, userDummy, users)
	fmt.Println(users)
}
