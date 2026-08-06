package models

import (
	"errors"
	"strings"
)

type User struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
}

func (u User) Validate() error {
	if strings.TrimSpace(u.Username) == "" || strings.TrimSpace(u.Password) == "" {
		return errors.New("username and password are required")
	}
	return nil
}
