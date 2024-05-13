package auth

import (
	"context"
)

type AuthService struct{}

func NewAuthService() (*AuthService, error) {
	return &AuthService{}, nil
}

func (auth *AuthService) ValidateToken(c context.Context, token string) (string, error) {
	return "something", nil
}
