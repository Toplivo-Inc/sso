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

	sS service.SessionService

	API   controllers.APIController
	OAuth controllers.OAuthController
	CRUD  controllers.CRUDController
}

func MustBuild() *Dependencies {
	var d Dependencies
	d.Config = config.MustLoad()
	d.DB = database.MustLoad(d.Config)

	// Set up repos
	uR := repository.NewUserRepo(d.DB)
	cR := repository.NewClientRepo(d.DB)
	aR := repository.NewAuthRepo(d.DB)
	sR := repository.NewScopeRepo(d.DB)

	// Set up services
	uS := service.NewUserService(uR, cR)
	aS := service.NewAuthService(uR)
	oaS := service.NewOAuthService(cR, aR, *d.Config)
	tS := service.NewTokenService()
	cS := service.NewClientService(cR, sR)
	d.sS = service.NewSessionService(uR)

	// Set up controllers
	d.API = controllers.NewAPI(aS, oaS, *d.Config)
	d.OAuth = controllers.NewOAuth(oaS, tS, uS, cS, *d.Config)
	d.CRUD = controllers.NewCRUD(uS, cS)

	return &d
}

func (d *Dependencies) SessionMiddleware() gin.HandlerFunc {
	return middlewares.SessionMiddleware(d.sS, *d.Config)
}
