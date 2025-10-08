package service

import (
	"sso/internal/storage"
	"sso/internal/storage/models"
)

type sessionService struct {
	userRepo storage.UserRepository
}

func NewSessionService(ur storage.UserRepository) SessionService {
	return &sessionService{ur}
}

// ValidateSession implements SessionService.
func (s *sessionService) ValidateSession(sessionToken string) (*models.Session, error) {
	return s.userRepo.GetSessionByToken(sessionToken)
}

// GenerateAccessToken implements SessionService.
func (s *sessionService) GenerateAccessToken() (string, error) {
	panic("unimplemented")
}

// GenerateIDToken implements SessionService.
func (s *sessionService) GenerateIDToken() (string, error) {
	panic("unimplemented")
}

// GenerateSessionToken implements SessionService.
func (s *sessionService) GenerateSessionToken() (string, error) {
	panic("unimplemented")
}
