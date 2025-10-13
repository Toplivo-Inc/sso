package repository

import (
	"sso/internal/models"
)

type ScopeRepository interface {
	Create(scope *models.Scope) error
	ScopeByID(id string) (*models.Scope, error)
	ClientScopes(clientID string) ([]models.Scope, error)
	Update(scope *models.Scope) error
	Delete(id string) error
}
