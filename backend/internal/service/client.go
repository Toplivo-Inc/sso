package service

import (
	"sso/internal/models"
	"sso/internal/repository"
	"sso/internal/utils"

	"github.com/google/uuid"
)

type ClientService interface {
	GetClients(paginated bool, limit, page int) []models.Client
	GetClientByID(id string) (*models.Client, error)
	AddClient(form *models.AddClientForm) (*models.Client, error)
	UpdateClient(clientID string, form models.UpdateClientForm) (*models.Client, error)
	DeleteClient(id string) error

	AddClientScope(clientID string, form *models.AddScopeForm) error
	Scopes(clientID string) []models.Scope
	UpdateScope(id string, form *models.UpdateScopeForm) (*models.Scope, error)
	DeleteScope(id string) error
}

type clientService struct {
	clientRepo repository.ClientRepository
	scopeRepo  repository.ScopeRepository
}

func NewClientService(cR repository.ClientRepository, sR repository.ScopeRepository) ClientService {
	return &clientService{cR, sR}
}

func (s *clientService) GetClientByID(id string) (*models.Client, error) {
	return s.clientRepo.ClientByID(id)
}

func (s *clientService) Scopes(clientID string) []models.Scope {
	scopes := make([]models.Scope, 0)
	s.scopeRepo.ClientScopes(clientID)
	return scopes
}

func (s *clientService) AddClient(form *models.AddClientForm) (*models.Client, error) {
	n := models.Client{
		Name:        form.Name,
		HomepageURL: form.HomepageURL,
		CallbackURL: form.CallbackURL,
		Secret:      utils.RandomString(16),
	}
	err := s.clientRepo.Create(&n)
	if err != nil {
		return nil, err
	}
	return s.clientRepo.ClientByID(n.ID.String())
}

func (s *clientService) AddClientScope(clientID string, form *models.AddScopeForm) error {
	id, err := uuid.Parse(clientID)
	if err != nil {
		return err
	}
	scope := models.Scope{
		ClientID: id,
		Resource: form.Resource,
		Action: form.Action,
		Name: form.Name,
	}
	scope.Description.Scan(form.Description)
	return s.scopeRepo.Create(&scope)
}

func (s *clientService) GetClients(paginated bool, limit, page int) []models.Client {
	if !paginated {
		return s.clientRepo.Clients()
	} else {
		return s.clientRepo.ClientsPaginated(limit, page)
	}
}

func (s *clientService) UpdateClient(id string, form models.UpdateClientForm) (*models.Client, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	newClient := &models.Client{
		ID:          uid,
		Name:        form.Name,
		HomepageURL: form.HomepageURL,
		CallbackURL: form.CallbackURL,
	}
	newClient.AvatarURL.Scan(&form.AvatarURL)
	newClient.Description.Scan(&form.Description)

	err = s.clientRepo.Update(newClient)
	if err != nil {
		return nil, err
	}
	return newClient, nil
}

func (s *clientService) DeleteClient(id string) error {
	return s.clientRepo.Delete(id)
}

func (s *clientService) UpdateScope(id string, form *models.UpdateScopeForm) (*models.Scope, error) {
	// FIXME: fuck what am i doing
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	newScope := &models.Scope{
		ID:          uid,
	}
	if form.Resource != nil {
		newScope.Resource = *form.Resource
	}
	if form.Action != nil {
		newScope.Action = *form.Action
	}
	if form.Name != nil {
		newScope.Name = *form.Name
	}
	newScope.Description.Scan(form.Description)

	err = s.scopeRepo.Update(newScope)
	if err != nil {
		return nil, err
	}
	return newScope, nil
}

func (s *clientService) DeleteScope(id string) error {
	return s.scopeRepo.Delete(id)
}
