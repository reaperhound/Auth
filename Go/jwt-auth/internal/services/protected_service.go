package services

type ProtectedService struct {
}

func NewProtectedService() *ProtectedService {
	return &ProtectedService{}
}

func (s *ProtectedService) Hey() (string, error) {
	message := "Hey this is a protected route"
	return message, nil
}
