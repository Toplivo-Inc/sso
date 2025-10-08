// Package service.
package service

import (
	"sso/internal/storage"
	"sso/internal/storage/models"
	"sso/internal/utils"
	"sso/pkg/errors"
)

type userService struct {
	userRepo   storage.UserRepository
	clientRepo storage.ClientRepository
}

func NewUseService(userRepo storage.UserRepository, clientRepo storage.ClientRepository) AuthService {
	return &userService{
		userRepo,
		clientRepo,
	}
}

// Register implements UserService.
func (s *userService) Register(form *models.UserRegisterForm) error {
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
func (s *userService) Login(form *models.UserLoginForm, metadata *models.LoginMetadata) (string, error) {
	found := false

	// TODO: create a new session with token based on metadata
	user, err := s.userRepo.GetUserByEmail(form.Login)
	if err == nil {
		if !utils.CheckPasswordHash(form.Password, user.PasswordHash.String) {
			return "", errors.AppErr(401, "incorrect password")
		}
		found = true
	}

	if !found {
		user, err = s.userRepo.GetUserByName(form.Login)
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

	createSession := models.Session{
		UserID:    user.ID,
		UserAgent: metadata.UserAgent,
		UserIP:    metadata.IP,
		// TODO: maybe some other generation
		SessionToken: utils.RandomString(32),
	}
	err = s.userRepo.CreateSession(&createSession)
	if err != nil {
		return "", err
	}
	session, err := s.userRepo.GetSessionByID(createSession.ID.String())
	if err != nil {
		return "", err
	}
	return session.SessionToken, nil
}

func (s *userService) FindUserByID(id string) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) FindUserPermissions(userID, clientID string) []models.Permission {
	return s.userRepo.GetPermissions(userID, clientID)
}
