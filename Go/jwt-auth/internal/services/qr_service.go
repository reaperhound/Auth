package services

import (
	"encoding/base64"

	"github.com/pquerna/otp"
	"github.com/skip2/go-qrcode"
)

type TwoFAData struct {
	QRCode string
	Secret string
}

func Generate2FA(key *otp.Key) (*TwoFAData, error) {
	qr, err := qrcode.Encode(
		key.URL(),
		qrcode.Medium,
		256,
	)

	if err != nil {
		return nil, err
	}

	return &TwoFAData{
		QRCode: base64.RawStdEncoding.EncodeToString(qr),
		Secret: key.Secret(),
	}, nil
}
