package storage

import (
	"sso/internal/storage/models"

	"gorm.io/gorm"
)


type sessionRepo struct {
	db *gorm.DB
}


func NewSessionRepo(db *gorm.DB) SessionRepository {
	return &sessionRepo{
		db: db,
	}
}

// Create inserts a new session based on provided model
func (r sessionRepo) Create(session *models.Session) error {
	result := r.db.Create(session)
	return result.Error
}

// GetByID selects a session with provided uuid
func (r sessionRepo) GetByID(uuid string) (*models.Session, error) {
	var session models.Session
	result := r.db.Where("id = ?", uuid).First(&session)

	return &session, result.Error
}

// FindByID selects a session with provided refresh token
func (r sessionRepo) GetByToken(sessionToken string) (*models.Session, error) {
	var session models.Session
	result := r.db.Where("session_token = ?", sessionToken).First(&session)

	return &session, result.Error
}

// GetByMetadata selects a session with provided userAgent and IP
func (r *sessionRepo) GetByMetadata(userAgent string, userIP string) (*models.Session, error) {
	var session models.Session
	result := r.db.Where("user_agent = ?", userAgent).Where("user_ip = ?", userIP).First(&session)

	return &session, result.Error
}

// Update updates a session based on provided model in place. It takes UUID from model.
func (r sessionRepo) Update(session *models.Session) error {
	result := r.db.Model(&session).Updates(session)
	return result.Error
}

// Delete deletes a session based on provided uuid.
func (r sessionRepo) Delete(uuid string) error {
	result := r.db.Delete(&models.Session{}, uuid)
	return result.Error
}
