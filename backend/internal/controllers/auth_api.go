// Package controllers defines HTTP handlers
package controllers

import (
	"sso/internal/service"
	"sso/internal/models"
	"sso/internal/config"
	"sso/internal/errors"

	"github.com/gin-gonic/gin"
)

type APIController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type api struct {
	authService  service.AuthService
	oauthService service.OAuthService
	config       config.Config
}

func NewAPI(as service.AuthService, os service.OAuthService, cfg config.Config) APIController {
	return &api{as, os, cfg}
}

// Register godoc
//
// @Summary register a new user
// @Description registers a new user
// @Tags JWT
// @accept application/x-www-form-urlencoded
// @Param form body models.UserRegisterForm true "Registration form"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /api/v1/register [post]
func (a api) Register(c *gin.Context) {
	var form models.UserRegisterForm
	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		c.Error(errors.AppErr(400, err.Error()))
		return
	}

	err := a.authService.Register(&form)
	if err != nil {
		c.Error(errors.AppErr(400, err.Error()))
		return
	}

	c.Status(201)
}

// Login godoc
//
// @Summary Login
// @Description returns a session token in httpOnly cookies
// @Tags JWT
// @accept json
// @Param form body models.UserLoginForm true "Login form"
// @Success 200 {string} access_token
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/v1/login [post]
func (a api) Login(c *gin.Context) {
	var form models.UserLoginForm
	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		c.Error(errors.AppErr(400, err.Error()))
		return
	}

	metadata := models.LoginMetadata{
		UserAgent: c.Request.UserAgent(),
		IP:        c.RemoteIP(),
	}

	sessionToken, err := a.authService.Login(&form, &metadata)
	if err != nil {
		c.Error(err)
		return
	}

	c.SetCookie("TOPLIVO_SESSION_TOKEN", sessionToken, 3600 * 24 * 365, "/", "localhost", a.config.App.Production, true)
	c.Status(201)
}
