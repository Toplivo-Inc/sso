// Package models
package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uuid.UUID      `json:"id" gorm:"primaryKey; type:uuid; default:gen_random_uuid()"`
	Username      string         `json:"username" gorm:"unique; not null"`
	Email         sql.NullString `json:"email" gorm:"unique"`
	EmailVerified bool           `json:"email_verified" gorm:"not null; default:false"`
	PasswordHash  sql.NullString
	AvatarURL     string       `json:"avatar_url"`
	LastLoginAt   *time.Time   `json:"last_login_at,omitempty"`
	IsBlocked     bool         `json:"is_blocked" gorm:"default:false; not null"`
	BlockedAt     *time.Time   `json:"blocked_at,omitempty"`
	Permissions   []Permission `gorm:"many2many:permission_to_user"`
	// GithubID     sql.NullInt32  `json:"github_id"`
	// TelegramID   sql.NullInt32  `json:"google_id,omitempty"`
}

type Session struct {
	gorm.Model
	ID     uuid.UUID `json:"id" gorm:"primaryKey; type:uuid; default:gen_random_uuid()"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid; foreignKey; references:users.id; constraint:OnDelete:CASCADE"`
	// UserAgent and UserIP are composite unique
	UserAgent    string `json:"user_agent" gorm:"not null; uniqueIndex:idx_user_agent"`
	UserIP       string `json:"user_ip" gorm:"type:inet; not null; uniqueIndex:idx_user_agent"`
	SessionToken string `json:"session_token" gorm:"not null"`
}

// Client is a registered app that is using TOPLIVO SSO for login.
type Client struct {
	gorm.Model
	ID          uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"not null; unique"`
	HomepageURL string         `json:"homepage_url" gorm:"not null; unique"`
	CallbackURL string         `json:"callback_url" gorm:"not null; unique"`
	Secret      string         `json:"secret" gorm:"not null"`
	AvatarURL   sql.NullString `json:"avatar_url"`
	Description sql.NullString `json:"description"`
}

type Permission struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"primaryKey; type:uuid; default:gen_random_uuid()"`
	ClientID uuid.UUID `json:"client_id" gorm:"type:uuid; foreignKey; references:clients.id; constraint:OnDelete:CASCADE"`
	// If permission is email:write, then email is resource and write is action
	Resource string `json:"resource" gorm:"not null"`
	Action   string `json:"action" gorm:"not null"`
	// Display name
	Name        string         `json:"name" gorm:"not null"`
	Description sql.NullString `json:"description"`
}

func (p Permission) ScopeString() string {
	return p.Resource + ":" + p.Action
}

type AuthRequest struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"primaryKey; type:uuid; default:gen_random_uuid()"`
	ResponseType        string    `gorm:"not null"`
	ClientID            uuid.UUID `gorm:"type:uuid; not null; foreignKey; references:clients.id; constraint:OnDelete:CASCADE"`
	RedirectURI         string    `gorm:"not null"`
	Scope               string    `gorm:"not null; default=openid+profile"`
	State               string
	Code                sql.NullString `gorm:"unique"`
	CodeChallenge       string
	CodeChallengeMethod string
	UserID              uuid.UUID `gorm:"type:uuid; foreignKey; references:users.id; constraint:OnDelete:CASCADE"` // Set after login
	ExpiresAt           time.Time `gorm:"not null"`
}
