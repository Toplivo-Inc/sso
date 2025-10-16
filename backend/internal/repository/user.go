// Package repository defines repositories over tables in DB
package repository

import (
	"sso/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	// user operations
	CreateUser(user *models.User) error
	Users() []models.User
	UsersPaginated(limit, page int) []models.User
	UserByID(id string) (*models.User, error)
	UserByEmail(email string) (*models.User, error)
	UserByName(name string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	SoftDeleteUser(id string) error
	RestoreUser(id string) error

	// session operations
	CreateSession(session *models.Session) error
	Sessions(id string) []models.Session
	SessionByID(id string) (*models.Session, error)
	SessionByToken(refreshToken string) (*models.Session, error)
	SessionByMetadata(userAgent string, userIP string) (*models.Session, error)
	UpdateSession(session *models.Session) error
	DeleteSession(id string) error

	// scope operations
	GetScopes(userID, clientID string) []models.Scope
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

// CreateUser inserts a new user based on provided model
func (r userRepo) CreateUser(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

// UserByID selects a user with provided uuid
func (r userRepo) UserByID(id string) (*models.User, error) {
	var user models.User
	result := r.db.Where("id = ?", id).First(&user)

	return &user, result.Error
}

// UserByEmail selects a user with provided uuid
func (r userRepo) UserByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)

	return &user, result.Error
}

// UserByName selects a user with provided uuid
func (r userRepo) UserByName(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)

	return &user, result.Error
}

// UpdateUser updates a user based on provided model in place. It takes UUID from model.
func (r userRepo) UpdateUser(user *models.User) error {
	result := r.db.Model(&user).Updates(user)
	return result.Error
}

// DeleteUser deletes a user based on provided uuid.
func (r userRepo) DeleteUser(id string) error {
	result := r.db.Where("id = ?", id).Unscoped().Delete(&models.User{})
	return result.Error
}

// SoftDeleteUser deletes a user based on provided uuid.
func (r userRepo) SoftDeleteUser(uuid string) error {
	result := r.db.Where("id = ?", uuid).Delete(&models.User{})
	return result.Error
}

// RestoreUser restores a soft deleted user
func (r userRepo) RestoreUser(id string) error {
	return r.db.Unscoped().Model(&models.User{}).Where("id = ?", id).Update("deleted_at", nil).Error
}

// GetScopes selects finds user scopes for provided client
func (r userRepo) GetScopes(userID, clientID string) []models.Scope {
	perms := make([]models.Scope, 0)

	r.db.Raw(`
	SELECT (p.*) FROM users AS u
		LEFT OUTER JOIN scope_to_user AS pu ON u.id = pu.user_id AND u.id = ?
		INNER JOIN scopes AS p ON p.id = pu.scope_id AND p.client_id = ?;`, userID, clientID).
		Scan(&perms)

	return perms
}

// CreateSession inserts a new session based on provided model
func (r userRepo) CreateSession(session *models.Session) error {
	result := r.db.Create(session)
	return result.Error
}

// SessionByID selects a session with provided uuid
func (r userRepo) SessionByID(id string) (*models.Session, error) {
	var session models.Session
	result := r.db.Where("id = ?", id).First(&session)

	return &session, result.Error
}

// SessionByToken selects a session with provided refresh token
func (r userRepo) SessionByToken(sessionToken string) (*models.Session, error) {
	var session models.Session
	result := r.db.Where("session_token = ?", sessionToken).First(&session)

	return &session, result.Error
}

// SessionByMetadata selects a session with provided userAgent and IP
func (r *userRepo) SessionByMetadata(userAgent string, userIP string) (*models.Session, error) {
	var session models.Session
	result := r.db.Where("user_agent = ?", userAgent).Where("user_ip = ?", userIP).First(&session)

	return &session, result.Error
}

// UpdateSession updates a session based on provided model in place. It takes UUID from model.
func (r userRepo) UpdateSession(session *models.Session) error {
	result := r.db.Model(&session).Updates(session)
	return result.Error
}

// DeleteSession deletes a session based on provided uuid.
func (r userRepo) DeleteSession(id string) error {
	result := r.db.Where("id = ?", id).Unscoped().Delete(&models.Session{})
	return result.Error
}

// Users selects all users
func (r *userRepo) Users() []models.User {
	users := make([]models.User, 0)
	r.db.Find(&users)
	return users
}

// Sessions selects all sessions
func (r *userRepo) Sessions(id string) []models.Session {
	sessions := make([]models.Session, 0)
	r.db.Where("user_id = ?", id).Find(&sessions)
	return sessions
}

// UsersPaginated selects a page of users
func (r *userRepo) UsersPaginated(limit int, page int) []models.User {
	users := make([]models.User, 0)
	r.db.Limit(limit).Offset(limit * (page - 1)).Find(&users)
	return users
}
