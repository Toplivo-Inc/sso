package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey; type:uuid; default:gen_random_uuid()"`
	Token     string    `json:"token" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	ExpiresAt time.Time `json:"expires_at" gorm:"default:CURRENT_TIMESTAMP + '365 days'"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid; foreignKey; references:users.id; constraint:OnDelete:CASCADE"`
	// UserAgent and UserIP are composite unique
	UserAgent string `json:"user_agent" gorm:"not null; uniqueIndex:idx_user_agent"`
	UserIP    string `json:"user_ip" gorm:"type:inet; not null; uniqueIndex:idx_user_agent"`
}

type SessionResponse struct {
	ID        string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
	UserID    string `json:"user_id"`
	// UserAgent and UserIP are composite unique
	UserAgent string `json:"user_agent"`
	UserIP    string `json:"user_ip"`
}

func (s Session) ToResponse() SessionResponse {
	return SessionResponse{
		ID: s.ID.String(),
		CreatedAt: s.CreatedAt,
		ExpiresAt: s.ExpiresAt,
		UserID: s.UserID.String(),
		UserAgent: s.UserAgent,
		UserIP: s.UserIP,
	}
}

func SessionsToResponses(in []Session) []SessionResponse {
	res := make([]SessionResponse, len(in))
	for i, el := range in {
		res[i] = el.ToResponse()
	}
	return res
}
