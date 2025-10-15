package service

import (
	"sso/internal/repository"
	"sso/internal/models"
)

type SessionService interface {
	ValidateSession(sessionToken string) (*models.Session, error)
	ValidateSessionWithMetadata(sessionToken, userAgent, userIP string) (*models.Session, error)

	// Token generation
	GenerateSessionToken() (string, error)
	GenerateAccessToken() (string, error)
	GenerateIDToken() (string, error)
}

type sessionService struct {
	userRepo repository.UserRepository
}

func NewSessionService(ur repository.UserRepository) SessionService {
	return &sessionService{ur}
}

// ValidateSession implements SessionService.
func (s *sessionService) ValidateSession(sessionToken string) (*models.Session, error) {
	return s.userRepo.SessionByToken(sessionToken)
}

// ValidateSessionWithMetadata implements SessionService.
func (s *sessionService) ValidateSessionWithMetadata(sessionToken, userAgent, userIP string) (*models.Session, error) {
	ses, err := s.userRepo.SessionByToken(sessionToken)
	if err != nil || ses.UserAgent != userAgent || ses.UserIP != userIP {
		return nil, err
	}
	return ses, nil
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
