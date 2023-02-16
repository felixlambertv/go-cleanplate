package user

import (
	"context"
	"fmt"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var service = NewUser()

func TestMain(m *testing.M) {
	fmt.Print("before")
	m.Run()
	fmt.Println("after")
}

func TestUser_GetUsers(t *testing.T) {
	users, err := service.GetUsers(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, &model.User{
		ID:    0,
		Name:  "",
		Email: "",
	}, users)
	fmt.Println(users)
}
