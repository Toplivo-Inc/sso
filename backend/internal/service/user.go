package service

import (
	"sso/internal/models"
	"sso/internal/repository"
	"sso/internal/utils"

	"github.com/google/uuid"
)

type UserService interface {
	GetUsers(paginated bool, limit, page int) []models.User
	GetUserByID(id string) (*models.User, error)
	GetUserScopes(userID, clientID string) []models.Scope
	UpdateUser(id string, form models.UpdateUserForm) (*models.User, error)
	DeleteUser(id string) error

	GetUserSessions(id string) []models.Session
	DeleteSession(id string) error
}

type userService struct {
	userRepo   repository.UserRepository
	clientRepo repository.ClientRepository
}

func NewUserService(userRepo repository.UserRepository, clientRepo repository.ClientRepository) UserService {
	return &userService{
		userRepo,
		clientRepo,
	}
}

func (s *userService) GetUsers(paginated bool, limit, page int) []models.User {
	if !paginated {
		return s.userRepo.Users()
	} else {
		return s.userRepo.UsersPaginated(limit, page)
	}
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.userRepo.UserByID(id)
}

func (s *userService) GetUserScopes(userID, clientID string) []models.Scope {
	return s.userRepo.GetScopes(userID, clientID)
}

func (s *userService) UpdateUser(id string, form models.UpdateUserForm) (*models.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		ID:        uid,
		Username:  form.Username,
		AvatarURL: form.AvatarURL,
		IsBlocked: form.IsBlocked,
	}
	if form.Email != "" {
		newUser.Email.Scan(&form.Email)
	}
	if form.Password != "" {
		hash, err := utils.HashPassword(form.Password)
		if err != nil {
			return nil, err
		}
		newUser.PasswordHash.Scan(&hash)
	}

	err = s.userRepo.UpdateUser(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepo.DeleteUser(id)
}

func (s *userService) GetUserSessions(id string) []models.Session {
	return s.userRepo.Sessions(id)
}

func (s *userService) DeleteSession(id string) error {
	return s.userRepo.DeleteSession(id)
}
