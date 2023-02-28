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

func NewUserRepo(db *gorm.DB, l logger.Interface) *UserRepo {
	return &UserRepo{db: db, l: l}
}

func (u *UserRepo) WithTrx(trxHandle *gorm.DB) repository.IUserRepo {
	if trxHandle == nil {
		u.l.Error("transaction db not found")
		return u
	}
	u.db = trxHandle
	return u
}

func (u *UserRepo) Store(user *model.User) (*model.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) FindAll() (users []model.User, err error) {
	result := u.db.Find(&users)

	if result.Error != nil {
		return users, err
	}

	return users, nil
}

func (u *UserRepo) FindByEmail(email string) (*model.User, error) {
	var user *model.User
	err := u.db.Where("email = ?", email).Take(&user).Error
	return user, err
}
