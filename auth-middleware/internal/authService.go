package internal

import (
	"context"
)

type AuthService struct{}

func (auth *AuthService) ValidateToken(c context.Context, token string) (string, error) {
	return "something", nil
}
