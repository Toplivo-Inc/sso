package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "sso/docs"
	"sso/internal/dependency"
	"sso/pkg/errors"
)

// @title Toplivo SSO API
// @version 1.0
// @host localhost:9100
func main() {
	dp := dependency.MustBuild()

	router := gin.New()
	router.Use(errors.ErrorHandling())
	router.Use(gin.Recovery())
	// FIXME: cors
	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
	}))

	api := router.Group("/api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	{
		v1 := api.Group("/v1")
		v1.POST("/register", dp.API.Register)
		v1.POST("/login", dp.API.Login)
	}
	{
		oauth := router.Group("/oauth")
		oauth.GET("/authorize", dp.SessionMiddleware(), dp.OAuth.Authorize)
		oauth.POST("/token", dp.OAuth.Token)
		oauth.GET("/userinfo", dp.OAuth.UserInfo)
		oauth.StaticFile("/jwks", "static/misc/jwks.json")
		router.StaticFile("/.well-known/openid-configuration", "static/misc/openid-configuration")
	}

	router.Static("/assets", "static/images")
	router.GET("/", dp.Frontend.Main)
	router.GET("/login", dp.Frontend.Login)
	router.GET("/register", dp.Frontend.Register)

	router.Run(":8080")
}
