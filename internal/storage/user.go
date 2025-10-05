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

// Create inserts a new user based on provided model
func (r userRepo) Create(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

// GetByID selects a user with provided uuid
func (r userRepo) GetByID(uuid string) (*models.User, error) {
	var user models.User
	result := r.db.Where("id = ?", uuid).First(&user)

	return &user, result.Error
}

// GetByEmail selects a user with provided uuid
func (r userRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)

	return &user, result.Error
}

// GetByName selects a user with provided uuid
func (r userRepo) GetByName(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)

	return &user, result.Error
}

// Update updates a user based on provided model in place. It takes UUID from model.
func (r userRepo) Update(user *models.User) error {
	result := r.db.Model(&user).Updates(user)
	return result.Error
}

// Delete deletes a user based on provided uuid.
func (r userRepo) Delete(uuid string) error {
	result := r.db.Delete(&models.User{}, uuid)
	return result.Error
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
