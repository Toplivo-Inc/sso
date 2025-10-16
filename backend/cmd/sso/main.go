package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "sso/docs"
	"sso/internal/dependency"
	"sso/internal/errors"
	m "sso/internal/middlewares"
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
		AllowOrigins:     []string{"http://localhost:9101", "https://localhost:9100"},
	}))

	{
		api := router.Group("/api")
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		v1 := api.Group("/v1")
		v1.POST("/register", dp.API.Register)
		v1.POST("/login", dp.API.Login)

		v1.GET("/users", m.Pagination(10), dp.CRUD.Users)
		v1.GET("/users/:id", dp.CRUD.UserByID)
		v1.GET("/users/:id/scopes/:second_id", dp.CRUD.UserScopes)
		v1.GET("/users/:id/sessions", dp.CRUD.UserSessions)
		v1.PUT("/users/:id", dp.CRUD.UpdateUser)
		v1.DELETE("/users/:id", dp.CRUD.DeleteUser)
		v1.DELETE("/sessions/:id", dp.CRUD.DeleteUserSession)

		v1.POST("/clients", dp.CRUD.AddClient)
		v1.GET("/clients", m.Pagination(10), dp.CRUD.Clients)
		v1.GET("/clients/:id", dp.CRUD.ClientByID)
		v1.PUT("/clients/:id", nil)
		v1.DELETE("/clients/:id", nil)

		v1.POST("/clients/:id/scopes", nil)
		v1.GET("/clients/:id/scopes", nil)
		v1.PUT("/clients/:id/scopes/:res/:action", nil)
		v1.DELETE("/clients/:id/scopes/:res/:action", nil)
	}
	{
		oauth := router.Group("/oauth")
		oauth.GET("/authorize", dp.SessionMiddleware(), dp.OAuth.Authorize)
		oauth.POST("/token", dp.OAuth.Token)
		oauth.GET("/userinfo", dp.SessionMiddleware(), dp.OAuth.UserInfo)
		oauth.GET("/logout", dp.SessionMiddleware(), dp.OAuth.Logout)
		oauth.StaticFile("/jwks", "static/misc/jwks.json")
		router.StaticFile("/.well-known/openid-configuration", "static/misc/openid-configuration")
	}

	router.Static("/assets", "static/images")
	router.GET("/login", func(c *gin.Context) {
		c.Redirect(302, fmt.Sprintf("http://localhost:9101/login?%s", c.Request.URL.RawQuery))
	})
	router.GET("/register", func(c *gin.Context) {
		c.Redirect(302, fmt.Sprintf("http://localhost:9101/register?%s", c.Request.URL.RawQuery))
	})

	router.Run(":8080")
}
