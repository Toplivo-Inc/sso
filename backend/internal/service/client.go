package service

import (
	"sso/internal/models"
	"sso/internal/repository"
	"sso/internal/utils"
)

type ClientService interface {
	GetClients(paginated bool, limit, page int) []models.Client
	GetClientByID(id string) (*models.Client, error)
	Permissions(clientID string, userID string) []models.Scope
	AddClient(form *models.AddClientForm) (*models.Client, error)
}

type clientService struct {
	clientRepo repository.ClientRepository
}

func NewClientService(clientRepo repository.ClientRepository) ClientService {
	return &clientService{clientRepo}
}

func (c *clientService) GetClientByID(id string) (*models.Client, error) {
	return c.clientRepo.ClientByID(id)
}

func (c *clientService) Permissions(clientID string, userID string) []models.Scope {
	scopes := make([]models.Scope, 0)
	return scopes
}

func (c *clientService) AddClient(form *models.AddClientForm) (*models.Client, error) {
	n := models.Client{
		Name:        form.Name,
		HomepageURL: form.HomepageURL,
		CallbackURL: form.CallbackURL,
		Secret:      utils.RandomString(16),
	}
	err := c.clientRepo.Create(&n)
	if err != nil {
		return nil, err
	}
	return c.clientRepo.ClientByID(n.ID.String())
}

func (s *clientService) GetClients(paginated bool, limit, page int) []models.Client {
	if !paginated {
		return s.clientRepo.Clients()
	} else {
		return s.clientRepo.ClientsPaginated(limit, page)
	}
}
