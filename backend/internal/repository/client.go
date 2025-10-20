package repository

import (
	"sso/internal/models"

	"gorm.io/gorm"
)

type ClientRepository interface {
	Create(client *models.Client) error
	ClientByID(id string) (*models.Client, error)
	ClientByName(name string) (*models.Client, error)
	Clients() []models.Client
	ClientsPaginated(limit int, page int) []models.Client
	Update(client *models.Client) error
	Delete(id string) error
}

type clientRepo struct {
	db *gorm.DB
}

func NewClientRepo(db *gorm.DB) ClientRepository {
	return &clientRepo{
		db: db,
	}
}

// Create inserts a new client based on provided model
func (r clientRepo) Create(client *models.Client) error {
	result := r.db.Create(client)
	return result.Error
}

// ClientByID selects an client with provided uuid
func (r clientRepo) ClientByID(id string) (*models.Client, error) {
	var client models.Client
	result := r.db.Where("id = ?", id).First(&client)

	return &client, result.Error
}

// ClientByName selects an client with provided name
func (r clientRepo) ClientByName(name string) (*models.Client, error) {
	var client models.Client
	result := r.db.Where("name = ?", name).First(&client)

	return &client, result.Error
}

// Update updates an client based on provided model in place. It takes UUID from model.
func (r clientRepo) Update(client *models.Client) error {
	result := r.db.Model(&client).Updates(client)
	return result.Error
}

// Delete deletes an client based on provided uuid.
func (r clientRepo) Delete(uuid string) error {
	result := r.db.Delete(&models.Client{}, uuid)
	return result.Error
}

// Clients implements ClientRepository.
func (r clientRepo) Clients() []models.Client {
	clients := make([]models.Client, 0)
	r.db.Find(&clients)
	return clients
}

// ClientsPaginated implements ClientRepository.
func (r clientRepo) ClientsPaginated(limit int, page int) []models.Client {
	clients := make([]models.Client, 0)
	r.db.Limit(limit).Offset(limit * (page - 1)).Find(&clients)
	return clients
}
