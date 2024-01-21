package service

import (
	"atm/internal/service/auth"
)

type Service struct {
	Auth auth.IAuthService
}

func NewService(AuthURL string) *Service {
	return &Service{
		Auth: auth.NewAuthService(AuthURL),
	}
}
