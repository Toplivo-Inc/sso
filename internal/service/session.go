package service

import (
	"sso/internal/storage"
	"sso/internal/storage/models"
)

type sessionService struct {
	sessionRepo storage.SessionRepository
	userRepo    storage.UserRepository
}

func NewSessionService(sr storage.SessionRepository, ur storage.UserRepository) SessionService {
	return &sessionService{sr, ur}
}

// ValidateSession implements SessionService.
func (s *sessionService) ValidateSession(sessionToken string) (*models.Session, error) {
	return s.sessionRepo.GetByToken(sessionToken)
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
