package repository

import (
	"sso/internal/models"

	"gorm.io/gorm"
)

type ScopeRepository interface {
	Create(scope *models.Scope) error
	ScopeByID(id string) (*models.Scope, error)
	ClientScopes(clientID string) []models.Scope
	Update(scope *models.Scope) error
	Delete(id string) error
}

type scopeRepo struct {
	db *gorm.DB
}

// ClientScopes implements ScopeRepository.
func (r *scopeRepo) ClientScopes(clientID string) []models.Scope {
	scopes := make([]models.Scope, 0)
	r.db.Joins("JOIN clients ON clients.id = ?", clientID).Find(&scopes)
	return scopes
}

// Create implements ScopeRepository.
func (r *scopeRepo) Create(scope *models.Scope) error {
	result := r.db.Create(scope)
	return result.Error
}

// Delete implements ScopeRepository.
func (r *scopeRepo) Delete(id string) error {
	result := r.db.Where("id = ?", id).Unscoped().Delete(&models.Scope{})
	return result.Error
}

// ScopeByID implements ScopeRepository.
func (r *scopeRepo) ScopeByID(id string) (*models.Scope, error) {
	var scope models.Scope
	result := r.db.Where("id = ?", id).First(&scope)

	return &scope, result.Error
}

// Update implements ScopeRepository.
func (r *scopeRepo) Update(scope *models.Scope) error {
	result := r.db.Model(&scope).Updates(scope)
	return result.Error
}

func NewScopeRepo(db *gorm.DB) ScopeRepository {
	return &scopeRepo{
		db: db,
	}
}
