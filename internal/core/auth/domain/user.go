package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID          string
	Email       string
	Username    string
	Password    string
	DisplayName string
	AvatarURL   string
	IsActive    bool
}

func NewUser(email, username, password string) (*User, error) {
	return &User{
		ID:          uuid.New().String(),
		Email:       email,
		Username:    username,
		Password:    password,
		DisplayName: username,
		IsActive:    true,
	}, nil
}
