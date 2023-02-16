package user

import (
	"context"
	"github.com/felixlambertv/go-cleanplate/internal/model"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUsers(ctx context.Context) (model.User, error) {
	return model.User{}, nil
}
