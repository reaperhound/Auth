package services

import "fmt"

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) SignUp() string {
	const msg string = "Signed up"
	fmt.Println(msg)
	return msg
}
