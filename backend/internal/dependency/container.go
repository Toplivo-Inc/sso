// Package dependency defines dependency injection container with config, database and controllers
package dependency

import (
	"sso/internal/config"
	"sso/internal/controllers"
	"sso/internal/database"
	"sso/internal/middlewares"
	"sso/internal/repository"
	"sso/internal/service"

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
	d.DB = database.MustLoad(d.Config)

	// Set up repos
	userRepo := repository.NewUserRepo(d.DB)
	clientRepo := repository.NewClientRepo(d.DB)
	authRepo := repository.NewAuthRepo(d.DB)

	// Set up services
	userService := service.NewAuthService(userRepo, clientRepo)
	oauthService := service.NewOAuthService(clientRepo, authRepo, *d.Config)
	clientService := service.NewClientService(clientRepo)
	d.sessionService = service.NewSessionService(userRepo)

	// Set up controllers
	d.API = controllers.NewAPI(userService, oauthService, *d.Config)
	d.OAuth = controllers.NewOAuth(oauthService, userService, clientService, *d.Config)

	return &d
}

func (d *Dependencies) SessionMiddleware() gin.HandlerFunc {
	return middlewares.SessionMiddleware(d.sessionService, *d.Config)
}
