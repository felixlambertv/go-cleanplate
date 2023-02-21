package auth

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/felixlambertv/go-cleanplate/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var userRepoMock = new(mocks.IUserRepo)
var authService = NewAuthService(userRepoMock)

var LoginRequest = request.LoginRequest{
	Email:    "user@test.com",
	Password: "password",
}

var ErrorLoginRequest = request.LoginRequest{
	Email:    "user@example.com",
	Password: "Password",
}

var userDummy = &model.User{
	ID:        0,
	Email:     "user@test.com",
	Password:  "password",
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeletedAt: gorm.DeletedAt{},
}

func TestMain(m *testing.M) {
	fmt.Print("before")

	hashedPassword, err := authService.EncryptPassword(userDummy.Password)
	if err != nil {
		log.Fatal("error on encrypting password", err)
	}

	userDummy.Password = hashedPassword

	m.Run()
	fmt.Println("after")
}

func TestAuth_Login(t *testing.T) {
	t.Setenv("APP_SECRET", "randomtestblabla")
	t.Setenv("TOKEN_HOUR_LIFESPAN", "1")

	userRepoMock.On("FindByEmail", userDummy.Email).Return(userDummy, nil)
	userRepoMock.On("FindByEmail", ErrorLoginRequest.Email).Return(nil, gorm.ErrRecordNotFound)
	user, token, err := authService.Login(LoginRequest)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, user, userDummy)
	assert.Equal(t, assert.NotNil(t, token), true)
	assert.Equal(t, err, nil)
}

func TestAuth_LoginShouldError(t *testing.T) {
	type errorTestCases struct {
		description   string
		input         request.LoginRequest
		expectedError string
	}

	for _, scenario := range []errorTestCases{
		{
			description:   "Email not found",
			input:         ErrorLoginRequest,
			expectedError: gorm.ErrRecordNotFound.Error(),
		},
		{
			description: "Invalid password",
			input: request.LoginRequest{
				Email:    userDummy.Email,
				Password: ErrorLoginRequest.Password,
			},
			expectedError: "password is incorrect",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			_, _, err := authService.Login(scenario.input)

			assert.Equal(t, scenario.expectedError, err.Error())
		})
	}
}
