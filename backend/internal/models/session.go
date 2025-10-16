package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID     uuid.UUID `json:"id" gorm:"primaryKey; type:uuid; default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid; foreignKey; references:users.id; constraint:OnDelete:CASCADE"`
	// UserAgent and UserIP are composite unique
	UserAgent    string `json:"user_agent" gorm:"not null; uniqueIndex:idx_user_agent"`
	UserIP       string `json:"user_ip" gorm:"type:inet; not null; uniqueIndex:idx_user_agent"`
	SessionToken string `json:"session_token" gorm:"not null"`
}
