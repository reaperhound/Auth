package services

import (
	"github.com/pquerna/otp/totp"
)

type TwoFAService struct {
	qrService *QRService
}

func NewTwoFASerive() *TwoFAService {
	return &TwoFAService{}
}

func (s *TwoFAService) GenerateEnrollment(username string) (*TwoFAData, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Go2FA",
		AccountName: username,
	})

	if err != nil {
		return nil, err
	}

	return s.qrService.GenerateQRCode(key)
}
