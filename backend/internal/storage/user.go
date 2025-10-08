package storage

import (
	"sso/internal/storage/models"

	"gorm.io/gorm"
)

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

// GetUserByID selects a user with provided uuid
func (r userRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	result := r.db.Where("id = ?", id).First(&user)

	return &user, result.Error
}

// GetUserByEmail selects a user with provided uuid
func (r userRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)

	return &user, result.Error
}

// GetUserByName selects a user with provided uuid
func (r userRepo) GetUserByName(username string) (*models.User, error) {
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

// FindByName selects a user with provided uuid
func (r userRepo) GetPermissions(userID, clientID string) []models.Permission {
	perms := make([]models.Permission, 0)

	r.db.Raw(`
	SELECT (p.*) FROM users AS u
		LEFT OUTER JOIN permission_to_user AS pu ON u.id = pu.user_id AND u.id = ?
		INNER JOIN permissions AS p ON p.id = pu.permission_id AND p.client_id = ?;`, userID, clientID).
		Scan(&perms)

	return perms
}

// CreateSession inserts a new session based on provided model
func (r userRepo) CreateSession(session *models.Session) error {
	result := r.db.Create(session)
	return result.Error
}

// GetSessionByID selects a session with provided uuid
func (r userRepo) GetSessionByID(id string) (*models.Session, error) {
	var session models.Session
	result := r.db.Where("id = ?", id).First(&session)

	return &session, result.Error
}

// FindByID selects a session with provided refresh token
func (r userRepo) GetSessionByToken(sessionToken string) (*models.Session, error) {
	var session models.Session
	result := r.db.Where("session_token = ?", sessionToken).First(&session)

	return &session, result.Error
}

// GetSessionByMetadata selects a session with provided userAgent and IP
func (r *userRepo) GetSessionByMetadata(userAgent string, userIP string) (*models.Session, error) {
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
	result := r.db.Delete(&models.Session{}, id)
	return result.Error
}
