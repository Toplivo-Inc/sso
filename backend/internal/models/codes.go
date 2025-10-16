// Package models
package models

import (
	"time"

	"github.com/google/uuid"
)

type AuthCodes struct {
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Code                string    `gorm:"primaryKey"`
	ResponseType        string    `gorm:"not null"`
	ClientID            uuid.UUID `gorm:"type:uuid; not null; foreignKey; references:clients.id; constraint:OnDelete:CASCADE"`
	RedirectURI         string    `gorm:"not null"`
	Scopes              string    `gorm:"not null; default=openid"`
	CodeChallenge       string
	CodeChallengeMethod string
	State               string
	UserID              uuid.UUID `gorm:"type:uuid; references:users.id; constraint:OnDelete:CASCADE"` // Set after login
	ExpiresAt           time.Time `gorm:"not null"`
	IsUsed              bool      `gorm:"not null; default=false"`
}
