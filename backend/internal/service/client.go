package service

import (
	"sso/internal/storage"
	"sso/internal/storage/models"
)


type clientService struct{
	clientRepo storage.ClientRepository
}

func NewClientService(clientRepo storage.ClientRepository) ClientService {
	return &clientService{clientRepo}
}

func (c *clientService) FindClientByID(id string) (*models.Client, error) {
	return c.clientRepo.ClientByID(id)
}

func (c *clientService) Permissions(clientID string, userID string) []models.Scope {
	permissions := make([]models.Scope, 0)
	return permissions
}
