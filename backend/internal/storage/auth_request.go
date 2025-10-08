package storage

import (
	"sso/internal/storage/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepository {
	return &authRepo{
		db: db,
	}
}

// Create inserts a new auth based on provided model
func (r authRepo) Create(req *models.AuthRequest) error {
	result := r.db.Create(req)
	return result.Error
}

// Create inserts a new auth based on provided model
func (r authRepo) CreateFromInput(input *models.AuthorizeInput) (*models.AuthRequest, error) {
	clientID, err := uuid.Parse(input.ClientID)
	if err != nil {
		return nil, err
	}
	req := models.AuthRequest{
		ResponseType:        string(input.ResponseType),
		ClientID:            clientID,
		RedirectURI:         input.RedirectURI,
		State:               input.State,
		CodeChallenge:       input.CodeChallenge,
		CodeChallengeMethod: string(input.CodeChallengeMethod),
	}
	result := r.db.Create(&req)
	return &req, result.Error
}

// GetByID selects an auth with provided uuid
func (r authRepo) GetByID(id string) (*models.AuthRequest, error) {
	var auth models.AuthRequest
	result := r.db.Where("id = ?", id).First(&auth)

	return &auth, result.Error
}

// GetByCode selects an auth with provided code
func (r authRepo) GetByCode(code string) (*models.AuthRequest, error) {
	var auth models.AuthRequest
	result := r.db.Where("code = ?", code).First(&auth)

	return &auth, result.Error
}


// Update updates an auth based on provided model in place. It takes UUID from model.
func (r authRepo) Update(req *models.AuthRequest) error {
	result := r.db.Model(&req).Updates(req)
	return result.Error
}

// Delete deletes an auth based on provided uuid.
func (r authRepo) Delete(uuid string) error {
	result := r.db.Delete(&models.AuthRequest{}, uuid)
	return result.Error
}
