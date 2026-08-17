package services

import (
	"fmt"
	"jwt-auth/internal/models"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type UserStore interface {
	ReadUsers() ([]models.User, error)
	WriteUsers([]models.User) error
}

type QRGenerator interface {
	GenerateQRCode(key *otp.Key) (*TwoFAData, error)
}

type TwoFAService struct {
	qrService   QRGenerator
	fileService UserStore
}

func NewTwoFASerive(qrService QRGenerator, fileService UserStore) *TwoFAService {
	return &TwoFAService{
		qrService:   qrService,
		fileService: fileService,
	}
}

func (s *TwoFAService) GenerateEnrollment(username string) (*TwoFAData, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Go2FA",
		AccountName: username,
	})

	if err != nil {
		return nil, err
	}

	users, err := s.fileService.ReadUsers()
	if err != nil {
		return nil, err
	}

	var foundUser *models.User

	for i, u := range users {
		if u.Username == username {
			foundUser = &users[i]
		}
	}

	if foundUser == nil {
		return nil, fmt.Errorf("user %q not found", username)
	}

	foundUser.TOTPSecret = key.Secret()

	err = s.fileService.WriteUsers(users)
	if err != nil {
		return nil, err
	}

	return s.qrService.GenerateQRCode(key)
}
