package repository

import (
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"gorm.io/gorm"
)

// IUser

type (
	IUserRepo interface {
		WithTrx(trxHandle *gorm.DB) IUserRepo
		FindAll() ([]model.User, error)
		Store(user *model.User) (*model.User, error)
	}
)
