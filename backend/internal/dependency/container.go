// Package dependency defines dependency injection container with config, database and controllers
package dependency

import (
	"sso/internal/controllers"
	"sso/internal/middlewares"
	"sso/internal/service"
	"sso/internal/storage"
	"sso/pkg/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Dependencies struct {
	Config *config.Config
	DB     *gorm.DB

	sessionService service.SessionService

	API      controllers.APIController
	OAuth    controllers.OAuthController
}

func MustBuild() *Dependencies {
	var d Dependencies
	d.Config = config.MustLoad()
	d.DB = storage.MustLoad(d.Config)

	// Set up repos
	userRepo := storage.NewUserRepo(d.DB)
	clientRepo := storage.NewClientRepo(d.DB)
	authRepo := storage.NewAuthRepo(d.DB)

	// Set up services
	userService := service.NewAuthService(userRepo, clientRepo)
	oauthService := service.NewOAuthService(clientRepo, authRepo)
	clientService := service.NewClientService(clientRepo)
	d.sessionService = service.NewSessionService(userRepo)

	// Set up controllers
	d.API = controllers.NewAPI(userService, oauthService, d.Config)
	d.OAuth = controllers.NewOAuth(oauthService, userService, clientService, d.Config)

	return &d
}

func (d *Dependencies) SessionMiddleware() gin.HandlerFunc {
	return middlewares.SessionMiddleware(d.sessionService, d.Config)
}
