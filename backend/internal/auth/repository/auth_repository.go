package repository

import (
	"errors"
)

type AuthRepository interface {
	GetUserByUsername(username string) (User, error)
}

type User struct {
	Username string
	Password string
}

type authRepo struct {
	users map[string]User
}

func NewAuthRepository() AuthRepository {
	// In a real app, this would be a DB connection. Here, it's just an in-memory map.
	return &authRepo{
		users: map[string]User{
			"admin": {Username: "admin", Password: "admin123"},
		},
	}
}

func (r *authRepo) GetUserByUsername(username string) (User, error) {
	if user, exists := r.users[username]; exists {
		return user, nil
	}
	return User{}, errors.New("user not found")
}
