package user

import (
	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/felixlambertv/go-cleanplate/internal/repository"
	"github.com/felixlambertv/go-cleanplate/internal/service"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repository.IUserRepo
}

func (u *UserService) CreateUser(req request.CreateUserRequest) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	user, err = u.userRepo.Store(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *UserService) GetUsers() ([]model.User, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserService(userRepo repository.IUserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) WithTrx(trxHandle *gorm.DB) service.IUserService {
	u.userRepo = u.userRepo.WithTrx(trxHandle)
	return u
}
