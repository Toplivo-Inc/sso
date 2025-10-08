// Package dependency defines dependency injection container with config, database and handlers
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

	API      handlers.APIController
	OAuth    handlers.OAuthController
	Frontend *handlers.FrontendController
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
	userService := service.NewUseService(userRepo, clientRepo)
	oauthService := service.NewOAuthService(clientRepo, authRepo)
	clientService := service.NewClientService(clientRepo)
	d.sessionService = service.NewSessionService(userRepo)

	// Set up handlers
	d.API = handlers.NewAPI(userService, oauthService, d.Config)
	d.OAuth = handlers.NewOAuth(oauthService, userService, clientService, d.Config)
	d.Frontend = handlers.MustLoadFrontend()

	return &d
}

func (d *Dependencies) SessionMiddleware() gin.HandlerFunc {
	return middlewares.SessionMiddleware(d.sessionService, d.Config)
}
