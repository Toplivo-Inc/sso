// Package utils is a bunch of stuff that i dont know where to put
package utils

import (
	"fmt"
	"time"

	"sso/internal/models"

	"github.com/golang-jwt/jwt"
)

// NewAccessToken takes app ID, user info and user scopes and returns a new
// signed JWT token that lasts 5 minutes
func NewAccessToken(client *models.Client, user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["aud"] = client.ID.String()
	claims["sub"] = user.ID.String()
	scopes := make([]string, 0)
	for _, perm := range user.Scopes {
		scopes = append(scopes, perm.ScopeString())
	}
	claims["scopes"] = scopes
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(5)).Unix()

	tokenString, err := token.SignedString([]byte(client.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// NewRefreshToken returns a new signed JWT token that lasts 30 days
// func NewRefreshToken(sessionID string, userID string, userAgent string, userIP string) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["session_id"] = sessionID
// 	claims["user_id"] = userID
// 	claims["user_agent"] = userAgent
// 	claims["user_ip"] = userIP
// 	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24*30)).Unix()
//
// 	tokenString, err := token.SignedString([]byte(config.Misc.SecretKey))
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return tokenString, nil
// }

// VerifyToken returns an error if token is invalid or expired
func VerifyToken(app *models.Client, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(app.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
