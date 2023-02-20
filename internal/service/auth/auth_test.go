package auth

import (
	"fmt"
	"testing"

	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/felixlambertv/go-cleanplate/mocks"
	"github.com/felixlambertv/go-cleanplate/pkg/utils"
	"gotest.tools/assert"
)

var userRepoMock = new(mocks.IUserRepo)
var authService = NewAuthService(userRepoMock)

var userDummy = &model.User{}
var tokenDummy = utils.TokenStruct{}

func TestMain(m *testing.M) {
	fmt.Print("before")

	m.Run()
	fmt.Println("after")
}

func TestAuth_Login(t *testing.T) {
	userRepoMock.On("FindByEmail", "user@example.com").Return(*userDummy, nil)
	user, token, err := authService.Login(request.LoginRequest{
		Email:    "user@example.com",
		Password: "password",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	assert.Equal(t, *userDummy, &user)
	assert.Equal(t, tokenDummy, token)
}