// Package service.
package service

import (
	"sso/internal/errors"
	"sso/internal/models"
	"sso/internal/repository"
	"sso/internal/utils"
)

type AuthService interface {
	Register(form *models.UserRegisterForm) error
	Login(form *models.UserLoginForm, metadata *models.LoginMetadata) (string, error)
	FindUserByID(id string) (*models.User, error)
	FindUserPermissions(userID, clientID string) []models.Scope
}

type authService struct {
	userRepo   repository.UserRepository
	clientRepo repository.ClientRepository
}

func NewAuthService(userRepo repository.UserRepository, clientRepo repository.ClientRepository) AuthService {
	return &authService{
		userRepo,
		clientRepo,
	}
}

// Register implements UserService.
func (s *authService) Register(form *models.UserRegisterForm) error {
	user := models.User{}
	user.Username = form.Username
	// NOTE: shouldn't return err
	_ = user.Email.Scan(form.Email)

	hash, err := utils.HashPassword(form.Password)
	if err != nil {
		return err
	}
	_ = user.PasswordHash.Scan(hash)

	return s.userRepo.CreateUser(&user)
}

// Login implements UserService.
// Searches for user with provided login as either email or username and checks if password is valid.
// Returns session token if everything is good.
func (s *authService) Login(form *models.UserLoginForm, metadata *models.LoginMetadata) (string, error) {
	found := false

	// TODO: create a new session with token based on metadata
	user, err := s.userRepo.UserByEmail(form.Login)
	if err == nil {
		if !utils.CheckPasswordHash(form.Password, user.PasswordHash.String) {
			return "", errors.AppErr(401, "incorrect password")
		}
		found = true
	}

	if !found {
		user, err = s.userRepo.UserByName(form.Login)
		if err == nil {
			if !utils.CheckPasswordHash(form.Password, user.PasswordHash.String) {
				return "", errors.AppErr(401, "incorrect password")
			}
			found = true
		}
	}

	if !found {
		return "", errors.AppErr(401, "user not found")
	}

	session, err := s.userRepo.SessionByMetadata(metadata.UserAgent, metadata.IP)
	if err != nil {
		createSession := models.Session{
			UserID:       user.ID,
			UserAgent:    metadata.UserAgent,
			UserIP:       metadata.IP,
			SessionToken: utils.RandomString(32),
		}
		err = s.userRepo.CreateSession(&createSession)
		if err != nil {
			return "", err
		}
		session, err = s.userRepo.SessionByID(createSession.ID.String())
		if err != nil {
			return "", err
		}
	}

	return session.SessionToken, nil
}

func (s *authService) FindUserByID(id string) (*models.User, error) {
	return s.userRepo.UserByID(id)
}

func (s *authService) FindUserPermissions(userID, clientID string) []models.Scope {
	return s.userRepo.GetScopes(userID, clientID)
}
