package service

import (
	"errors"

	"github.com/javy99/MoneyWise/backend/internal/auth/repository"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return AuthService{repo: repo}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}
	// Here, you would generate a token (JWT) or any other auth mechanism
	return "fake-jwt-token", nil
}
