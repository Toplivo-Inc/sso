package service

import (
	"sso/internal/repository"
	"sso/internal/models"
)

type ClientService interface {
	FindClientByID(id string) (*models.Client, error)
	Permissions(clientID string, userID string) []models.Scope
}

type clientService struct{
	clientRepo repository.ClientRepository
}

func NewClientService(clientRepo repository.ClientRepository) ClientService {
	return &clientService{clientRepo}
}

func (c *clientService) FindClientByID(id string) (*models.Client, error) {
	return c.clientRepo.ClientByID(id)
}

func (c *clientService) Permissions(clientID string, userID string) []models.Scope {
	scopes := make([]models.Scope, 0)
	return scopes
}
