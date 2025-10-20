package models

import (
	"database/sql"

	"sso/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

type ClientResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	HomepageURL string  `json:"homepage_url"`
	CallbackURL string  `json:"callback_url"`
	Secret      string  `json:"secret"`
	AvatarURL   *string `json:"avatar_url"`
	Description *string `json:"description"`
}

func (c Client) ToResponse() ClientResponse {
	return ClientResponse{
		ID:          c.ID.String(),
		Name:        c.Name,
		HomepageURL: c.HomepageURL,
		CallbackURL: c.CallbackURL,
		Secret:      c.Secret,
		AvatarURL:   utils.ResolveNullString(c.AvatarURL),
		Description: utils.ResolveNullString(c.Description),
	}
}

func ClientsToResponses(cs []Client) []ClientResponse {
	res := make([]ClientResponse, len(cs))
	for i, c := range cs {
		res[i] = c.ToResponse()
	}
	return res
}

type AddClientForm struct {
	Name        string `json:"name" example:"my-app" binding:"required"`
	HomepageURL string `json:"homepage_url" example:"http://localhost:9102" binding:"required"`
	CallbackURL string `json:"callback_url" example:"http://localhost:9102/callback" binding:"required"`
	Description string `json:"description" example:"This is the greatest app of all time"`
}

type UpdateClientForm struct {
	Name        string `json:"name" example:"my-cool-app"`
	HomepageURL string `json:"homepage_url" example:"https://coolapp.com"`
	CallbackURL string `json:"callback_url" example:"https://coolapp.com/callback"`
	AvatarURL   string `json:"avatar_url" example:"https://pics.com/coolapp"`
	Description string `json:"description" example:"Coolest app eva"`
}
