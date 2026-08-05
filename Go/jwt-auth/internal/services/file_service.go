package services

import (
	"encoding/json"
	"jwt-auth/internal/models"
	"os"
)

type FileService struct{}

const fileName string = "users.json"

func NewFileService() *FileService {
	return &FileService{}
}

func (f *FileService) ReadUsers() ([]models.User, error) {
	var users []models.User

	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return users, nil //no file yet = no users, not an error
		}
		return nil, err
	}

	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (f *FileService) WriteUsers(users []models.User) error {
	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
