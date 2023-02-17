package user

import (
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/felixlambertv/go-cleanplate/internal/repository"
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"gorm.io/gorm"
)

type UserRepo struct {
	l  logger.Interface
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) WithTrx(trxHandle *gorm.DB) repository.IUserRepo {
	if trxHandle == nil {
		u.l.Error("transaction db not found")
		return u
	}
	u.db = trxHandle
	return u
}

func (u *UserRepo) Store() (*model.User, error) {
	return &model.User{
		ID:       0,
		Name:     "",
		Email:    "",
		Password: "",
	}, nil
}

func (u *UserRepo) FindAll() ([]model.User, error) {
	return []model.User{
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
	}, nil
}
