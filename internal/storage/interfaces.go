package storage

import "sso/internal/storage/models"

type UserRepository interface {
	// user operations
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByName(name string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	// session operations
	// permission operations
	GetPermissions(userID, clientID string) []models.Permission
}

type SessionRepository interface {
	Create(session *models.Session) error
	GetByID(id string) (*models.Session, error)
	GetByToken(refreshToken string) (*models.Session, error)
	GetByMetadata(userAgent string, userIP string) (*models.Session, error)
	Update(session *models.Session) error
	Delete(id string) error
}

type AuthRepository interface {
	Create(req *models.AuthRequest) error
	CreateFromInput(req *models.AuthorizeInput) (*models.AuthRequest, error)
	GetByID(id string) (*models.AuthRequest, error)
	GetByCode(code string) (*models.AuthRequest, error)
	Update(req *models.AuthRequest) error
	Delete(id string) error
}

type ClientRepository interface {
	Create(client *models.Client) error
	GetByID(id string) (*models.Client, error)
	GetByName(name string) (*models.Client, error)
	Update(client *models.Client) error
	Delete(id string) error
}

type PermissionRepository interface {
	Create(permission *models.Permission) error
	GetByID(id string) (*models.Permission, error)
	GetClientPermissions(clientID string) ([]models.Permission, error)
	Update(permission *models.Permission) error
	Delete(id string) error
}
