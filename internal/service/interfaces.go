package service

import (
	"context"
	"github.com/felixlambertv/go-cleanplate/internal/model"
)

// User Interface
type (
	User interface {
		GetUsers(ctx context.Context) (model.User, error)
	}

	UserRepo interface {
	}
)
