package auth

import (
	"atm/internal/models"
	"atm/internal/repository/auth"
)

type AuthService struct {
	repo auth.IAuthRepo
}

type IAuthService interface {
	CreateNewUser(user models.User) error
	GetUserByUsername(username string) (models.User, error)
	DeleteUserByUsername(username string) error
	CheckUserCreds(creds models.User) error
}

func NewAuthService(repo auth.IAuthRepo) *AuthService {
	return &AuthService{repo: repo}
}
