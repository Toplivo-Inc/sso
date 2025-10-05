package service

import "sso/internal/storage/models"

type AuthService interface {
	Register(form *models.UserRegisterForm) error
	Login(form *models.UserLoginForm, metadata *models.LoginMetadata) (string, error)
	FindUserByID(id string) (*models.User, error)
	FindUserPermissions(userID, clientID string) []models.Permission
}

type ClientService interface {
	FindClientByID(id string) (*models.Client, error)
	Permissions(clientID string, userID string) []models.Permission
}

type OAuthService interface {
	ValidateAuthorizeInput(models.AuthorizeInput) error
	NewAuthReq(models.AuthorizeInput) (*models.AuthRequest, error)
	FindAuthReq(id string) (*models.AuthRequest, error)
	FindAuthReqByCode(id string) (*models.AuthRequest, error)
	UpdateAuthReq(*models.AuthRequest) error
}

type SessionService interface {
	ValidateSession(sessionToken string) (*models.Session, error)

	// Token generation
	GenerateSessionToken() (string, error)
	GenerateAccessToken() (string, error)
	GenerateIDToken() (string, error)
}
