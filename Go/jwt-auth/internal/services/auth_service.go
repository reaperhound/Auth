package services

import (
	"errors"
	"fmt"
	"jwt-auth/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	fileService *FileService
}

func NewAuthService(fileService *FileService) *AuthService {
	return &AuthService{fileService: fileService}
}

func (s *AuthService) SignUp(username, password string) error {
	users, err := s.fileService.ReadUsers()

	if err != nil {
		return fmt.Errorf("failed to read users: %w", err)
	}

	for _, u := range users {
		if u.Username == username {
			return errors.New("username already exists")
		}
	}

	hashedPass, err := hashPassword(password)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	users = append(users, models.User{
		Username: username,
		Password: hashedPass,
	})

	if err := s.fileService.WriteUsers(users); err != nil {
		return fmt.Errorf("failed to signup user %w", err)
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hashedPass), nil
}

func (s *AuthService) Login(username, password string) error {
	users, err := s.fileService.ReadUsers()
	if err != nil {
		return fmt.Errorf("failed to read users: %w", err)
	}

	for _, u := range users {
		if u.Username == username {
			err := bcrypt.CompareHashAndPassword(
				[]byte(u.Password),
				[]byte(password),
			)
			if err != nil {
				return errors.New("Invalid credentials")
			}

			fmt.Println("Logged in")
			return nil
		}
	}

	return errors.New("invalid credentials")
}
