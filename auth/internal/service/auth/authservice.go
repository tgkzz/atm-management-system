package auth

import (
	auth2 "auth/internal/models"
	"auth/internal/repository/auth"
)

type AuthService struct {
	repo auth.IAuthRepo
}

type IAuthService interface {
	CreateNewUser(user auth2.User) error
	GetUserByUsername(username string) (auth2.User, error)
	DeleteUserByUsername(username string) error
	CheckUserCreds(creds auth2.User) (auth2.User, error)
	//Login(user models.User) (string, error)
}

func NewAuthService(repo auth.IAuthRepo) *AuthService {
	return &AuthService{repo: repo}
}
