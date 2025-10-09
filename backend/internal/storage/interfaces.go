package storage

import (
	"sso/internal/storage/models"
)

type UserRepository interface {
	// user operations
	CreateUser(user *models.User) error
	UserByID(id string) (*models.User, error)
	UserByEmail(email string) (*models.User, error)
	UserByName(name string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	SoftDeleteUser(id string) error
	RestoreUser(id string) error

	// session operations
	CreateSession(session *models.Session) error
	SessionByID(id string) (*models.Session, error)
	SessionByToken(refreshToken string) (*models.Session, error)
	SessionByMetadata(userAgent string, userIP string) (*models.Session, error)
	UpdateSession(session *models.Session) error
	DeleteSession(id string) error

	// permission operations
	GetPermissions(userID, clientID string) []models.Scope
}

type AuthRepository interface {
	Create(req *models.AuthRequest) error
	CreateFromInput(req *models.AuthorizeInput) (*models.AuthRequest, error)
	AuthReqByID(id string) (*models.AuthRequest, error)
	AuthReqByCode(code string) (*models.AuthRequest, error)
	Update(req *models.AuthRequest) error
	Delete(id string) error
}

type ClientRepository interface {
	Create(client *models.Client) error
	ClientByID(id string) (*models.Client, error)
	ClientByName(name string) (*models.Client, error)
	Update(client *models.Client) error
	Delete(id string) error
}

type ScopeRepository interface {
	Create(permission *models.Scope) error
	ScopeByID(id string) (*models.Scope, error)
	ClientScopes(clientID string) ([]models.Scope, error)
	Update(permission *models.Scope) error
	Delete(id string) error
}
