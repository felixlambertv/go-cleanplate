package service

import (
	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/felixlambertv/go-cleanplate/pkg/utils"
	"gorm.io/gorm"
)

// IUserService Interface
type (
	IUserService interface {
		WithTrx(trxHandle *gorm.DB) IUserService
		CreateUser(req request.CreateUserRequest) (*model.User, error)
		GetUsers() ([]model.User, error)
	}

	IAuthService interface {
		Login(req request.LoginRequest) (*model.User, *utils.Token, error)
		EncryptPassword(password string) (string, error)
	}
)
