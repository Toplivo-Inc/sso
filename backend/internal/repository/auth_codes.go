package repository

import (
	"sso/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(req *models.AuthCodes) error
	CreateFromInput(req *models.AuthorizeInput) (*models.AuthCodes, error)
	AuthReqByState(state string) (*models.AuthCodes, error)
	AuthReqByCode(code string) (*models.AuthCodes, error)
	Update(req *models.AuthCodes) error
	Delete(id string) error
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepository {
	return &authRepo{
		db: db,
	}
}

// Create inserts a new auth based on provided model
func (r authRepo) Create(req *models.AuthCodes) error {
	result := r.db.Create(req)
	return result.Error
}

// Create inserts a new auth based on provided model
func (r authRepo) CreateFromInput(input *models.AuthorizeInput) (*models.AuthCodes, error) {
	clientID, err := uuid.Parse(input.ClientID)
	if err != nil {
		return nil, err
	}
	req := models.AuthCodes{
		State:               input.State,
		ResponseType:        string(input.ResponseType),
		ClientID:            clientID,
		RedirectURI:         input.RedirectURI,
		CodeChallenge:       input.CodeChallenge,
		CodeChallengeMethod: string(input.CodeChallengeMethod),
	}
	result := r.db.Create(&req)
	return &req, result.Error
}

// AuthReqByID selects an auth with provided state
func (r authRepo) AuthReqByState(state string) (*models.AuthCodes, error) {
	var auth models.AuthCodes
	result := r.db.Where("state = ?", state).First(&auth)

	return &auth, result.Error
}

// AuthReqByCode selects an auth with provided code
func (r authRepo) AuthReqByCode(code string) (*models.AuthCodes, error) {
	var auth models.AuthCodes
	result := r.db.Where("code = ?", code).First(&auth)

	return &auth, result.Error
}


// Update updates an auth based on provided model in place. It takes UUID from model.
func (r authRepo) Update(req *models.AuthCodes) error {
	result := r.db.Model(&req).Updates(req)
	return result.Error
}

// Delete deletes an auth based on provided uuid.
func (r authRepo) Delete(uuid string) error {
	result := r.db.Delete(&models.AuthCodes{}, uuid)
	return result.Error
}
