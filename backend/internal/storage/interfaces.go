package storage

import (
	"sso/internal/storage/models"
)

type UserRepository interface {
	// user operations
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByName(name string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	SoftDeleteUser(id string) error
	RestoreUser(id string) error

	// session operations
	CreateSession(session *models.Session) error
	GetSessionByID(id string) (*models.Session, error)
	GetSessionByToken(refreshToken string) (*models.Session, error)
	GetSessionByMetadata(userAgent string, userIP string) (*models.Session, error)
	UpdateSession(session *models.Session) error
	DeleteSession(id string) error

	// permission operations
	GetPermissions(userID, clientID string) []models.Permission
}

type SessionRepository interface {
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
