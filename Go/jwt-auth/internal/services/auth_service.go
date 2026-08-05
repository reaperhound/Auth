package services

import (
	"encoding/json"
	"fmt"
	"jwt-auth/internal/models"
	"os"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) SignUp(username, password string) error {
	user := models.User{
		Username: username,
		Password: password,
	}

	err := s.appendUser(user)
	if err != nil {
		return fmt.Errorf("failed to sign up user: %w", err)
	}

	fmt.Println("Signed up")
	return nil
}

func (s *AuthService) appendUser(user models.User) error {
	const fileName = "users.json"

	var users []models.User

	// Read
	data, err := os.ReadFile(fileName)
	if err == nil {
		json.Unmarshal(data, &users)
	}

	// Add new user
	users = append(users, user)

	// Convert back to json
	updatedData, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, updatedData, 0644)
}
