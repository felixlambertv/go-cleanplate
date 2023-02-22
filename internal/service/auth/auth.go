package auth

import (
	"errors"

	"github.com/felixlambertv/go-cleanplate/config"
	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/felixlambertv/go-cleanplate/internal/repository"
	"github.com/felixlambertv/go-cleanplate/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	cfg      *config.Config
	userRepo repository.IUserRepo
}

func (a *AuthService) Login(req request.LoginRequest) (*model.User, *utils.Token, error) {
	user, err := a.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, nil, err
	}

	err = verifyPassword(user, req.Password)
	if err != nil {
		return nil, nil, err
	}

	token, err := utils.GenerateToken(user, a.cfg.App.TokenLifespan, a.cfg.App.Secret)
	if err != nil {
		return nil, nil, err
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

func NewAuthService(userRepo repository.IUserRepo, cfg *config.Config) *AuthService {
	return &AuthService{userRepo: userRepo, cfg: cfg}
}
