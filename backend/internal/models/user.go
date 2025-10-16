package models

import (
	"database/sql"
	"time"

	"sso/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRegisterForm is sent to /register handler
type UserRegisterForm struct {
	Username string `json:"username" example:"John" binding:"required"`
	Email    string `json:"email" example:"john@gmail.com" binding:"required,email"`
	Password string `json:"password" example:"not12345" binding:"required"`
}

// UserLoginForm is sent to /login handler. Login could be either username or email
type UserLoginForm struct {
	Login    string `json:"login" example:"John" binding:"required"`
	Password string `json:"password" example:"not12345" binding:"required"`
}

type User struct {
	gorm.Model
	ID            uuid.UUID      `gorm:"primaryKey; type:uuid; default:gen_random_uuid()"`
	Username      string         `gorm:"unique; not null"`
	Email         sql.NullString `gorm:"unique"`
	EmailVerified bool           `gorm:"not null; default:false"`
	PasswordHash  sql.NullString
	AvatarURL     string
	LastLoginAt   sql.NullTime
	IsBlocked     bool `gorm:"default:false; not null"`
	BlockedAt     sql.NullTime
	Scopes        []Scope `gorm:"many2many:scope_to_user"`
	// GithubID     sql.NullInt32  `json:"github_id"`
	// TelegramID   sql.NullInt32  `json:"google_id,omitempty"`
}

type UserResponse struct {
	ID            string     `json:"id"`
	Username      string     `json:"username"`
	Email         *string    `json:"email"`
	EmailVerified bool       `json:"email_verified"`
	AvatarURL     string     `json:"avatar_url"`
	LastLoginAt   *time.Time `json:"last_login_at"`
	IsBlocked     bool       `json:"is_blocked"`
	BlockedAt     *time.Time `json:"blocked_at"`
}

func (u User) ToResponse() UserResponse {
	return UserResponse{
		ID:            u.ID.String(),
		Username:      u.Username,
		Email:         utils.ResolveNullString(u.Email),
		EmailVerified: u.EmailVerified,
		AvatarURL:     u.AvatarURL,
		LastLoginAt:   utils.ResolveNullTime(u.LastLoginAt),
		IsBlocked:     u.IsBlocked,
		BlockedAt:     utils.ResolveNullTime(u.BlockedAt),
	}
}

func UsersToResponses(us []User) []UserResponse {
	res := make([]UserResponse, len(us))
	for i, u := range us {
		res[i] = u.ToResponse()
	}
	return res
}

type UpdateUserForm struct {
	Username  string `json:"username" example:"John"`
	// FIXME: no email binding
	Email     string `json:"email" example:"john@gmail.com"`
	Password  string `json:"password" example:"not12345"`
	AvatarURL string `json:"avatar_url" example:"http://avatars.com/1"`
	IsBlocked bool   `json:"is_blocked" example:"false"`
}
