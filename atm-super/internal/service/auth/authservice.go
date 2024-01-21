package auth

import (
	"atm/internal/models"
)

type AuthService struct {
	AuthURL string
}

type IAuthService interface {
	CreateNewUser(user models.User) (map[string]string, error)
}

func NewAuthService(AuthURL string) *AuthService {
	return &AuthService{AuthURL: AuthURL}
}
