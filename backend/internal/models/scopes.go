package models

import (
	"database/sql"
	"sso/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Scope struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"primaryKey; type:uuid; default:gen_random_uuid()"`
	ClientID uuid.UUID `json:"client_id" gorm:"type:uuid; foreignKey; references:clients.id; constraint:OnDelete:CASCADE"`
	// If scope is email:write, then email is resource and write is action
	Resource string `json:"resource" gorm:"not null"`
	Action   string `json:"action" gorm:"not null"`
	// Display name
	Name        string         `json:"name" gorm:"not null"`
	Description sql.NullString `json:"description"`
}

func (p Scope) String() string {
	return p.Resource + ":" + p.Action
}

type ScopeResponse struct {
	ID          string  `json:"id"`
	ClientID    string  `json:"client_id"`
	Resource    string  `json:"resource"`
	Action      string  `json:"action"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func ScopesToResponses(in []Scope) []ScopeResponse {
	res := make([]ScopeResponse, len(in))

	for i, s := range in {
		res[i] = ScopeResponse{
			ID:          s.ID.String(),
			ClientID:    s.ClientID.String(),
			Resource:    s.Resource,
			Action:      s.Action,
			Name:        s.Name,
			Description: utils.ResolveNullString(s.Description),
		}
	}
	return res
}
