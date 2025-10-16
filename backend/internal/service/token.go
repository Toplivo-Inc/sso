// Package utils is a bunch of stuff that i dont know where to put
package service

import (
	"fmt"
	"time"

	"sso/internal/config"
	"sso/internal/models"

	"github.com/golang-jwt/jwt"
)

type TokenService interface {
	NewAccessToken(client *models.Client, user *models.User, c config.Config) (string, error)
	NewIDToken(client *models.Client, user *models.User, c config.Config) (string, error)
	VerifyToken(client models.Client, tokenString string) (*jwt.Token, error)
}

type tokenService struct{}

func NewTokenService() TokenService {
	return tokenService{}
}

// NewAccessToken returns a signed JWT access token that lasts 5 minutes
func (t tokenService) NewAccessToken(client *models.Client, user *models.User, c config.Config) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = c.App.BaseURL
	claims["aud"] = client.ID.String()
	claims["sub"] = user.ID.String()
	iat := time.Now()
	claims["iat"] = iat.Unix()
	claims["exp"] = iat.Add(time.Minute * time.Duration(5)).Unix()

	scopes := make([]string, 0)
	for _, perm := range user.Scopes {
		scopes = append(scopes, perm.String())
	}
	claims["scopes"] = scopes

	tokenString, err := token.SignedString([]byte(client.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// NewIDToken returns JWT ID token that lasts 5 minutes
func (t tokenService) NewIDToken(client *models.Client, user *models.User, c config.Config) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = c.App.BaseURL
	claims["aud"] = client.ID.String()
	claims["sub"] = user.ID.String()
	iat := time.Now()
	claims["iat"] = iat.Unix()
	claims["exp"] = iat.Add(time.Minute * time.Duration(5)).Unix()

	tokenString, err := token.SignedString([]byte(client.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken returns an error if token is invalid or expired
func (t tokenService) VerifyToken(client models.Client, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(client.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
