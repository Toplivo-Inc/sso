package storage

import (
	"sso/internal/storage/models"

	"gorm.io/gorm"
)

type clientRepo struct {
	db *gorm.DB
}

func NewClientRepo(db *gorm.DB) ClientRepository {
	return &clientRepo {
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

// FindByID selects an client with provided name
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
