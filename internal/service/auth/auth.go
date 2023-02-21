package auth

import (
	"errors"

	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/felixlambertv/go-cleanplate/internal/repository"
	"github.com/felixlambertv/go-cleanplate/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.IUserRepo
}

func (a *AuthService) Login(req request.LoginRequest) (*model.User, utils.TokenStruct, error) {
	user, err := a.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, utils.TokenStruct{}, err
	}

	err = verifyPassword(user, req.Password)
	if err != nil {
		return nil, utils.TokenStruct{}, err
	}

	token, err := utils.GenerateToken(*user)
	if err != nil {
		return nil, utils.TokenStruct{}, err
	}

	return user, token, nil
}

func verifyPassword(u *model.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return errors.New("password is incorrect")
		default:
			return err
		}
	}

	return err
}

func (a *AuthService) EncryptPassword(password string) (string, error) {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}

func NewAuthService(userRepo repository.IUserRepo) *AuthService {
	return &AuthService{userRepo: userRepo}
}
