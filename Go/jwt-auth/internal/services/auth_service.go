package services

import (
	"errors"
	"fmt"
	"jwt-auth/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	fileService *FileService
	jwtService  *JwtServie
}

func NewAuthService(fileService *FileService, jwtService *JwtServie) *AuthService {
	return &AuthService{fileService: fileService, jwtService: jwtService}
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

func (s *AuthService) Login(username, password string) (string, string, error) {
	users, err := s.fileService.ReadUsers()
	if err != nil {
		return "", "", fmt.Errorf("failed to read users: %w", err)
	}

	// Find user and verify password
	var foundUser *models.User

	for _, u := range users {
		if u.Username == username {
			foundUser = &u
			break
		}
	}

	fmt.Println(*foundUser)

	if foundUser == nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := s.jwtService.GenerateAccessTok(username)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate accessToken: %w", err)
	}

	refreshToken, err := s.jwtService.GenerateRefreshTok(username)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refreshToken: %w", err)
	}

	fmt.Println("Logged in")
	return accessToken, refreshToken, nil
}
