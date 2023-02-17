package service

import (
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"gorm.io/gorm"
)

// IUserService Interface
type (
	IUserService interface {
		WithTrx(trxHandle *gorm.DB) IUserService
		GetUsers() ([]model.User, error)
	}
)
