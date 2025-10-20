package service

import (
	"sso/internal/models"
	"sso/internal/repository"
)

type SessionService interface {
	ValidateSession(sessionToken string) (*models.Session, error)
	ValidateSessionWithMetadata(sessionToken, userAgent, userIP string) (*models.Session, error)
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
