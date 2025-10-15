package middlewares

import (
	"sso/internal/config"
	"sso/internal/models"
	"sso/internal/service"

	"github.com/gin-gonic/gin"
)

// SessionMiddleware checks whether TOPLIVO_SESSION_TOKEN is provided and valid
func SessionMiddleware(sessionService service.SessionService, config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, err := c.Cookie("TOPLIVO_SESSION_TOKEN")
		if err != nil {
			// No session cookie, user is not authenticated
			c.Next()
			return
		}

		metadata := models.LoginMetadata{
			UserAgent: c.Request.UserAgent(),
			IP:        c.RemoteIP(),
		}

		session, err := sessionService.ValidateSessionWithMetadata(sessionToken, metadata.UserAgent, metadata.IP)
		if err != nil {
			// Invalid session, clear the cookie
			c.SetCookie("TOPLIVO_SESSION_TOKEN", "", -1, "/", "localhost", config.App.Production, true)
			c.SetCookie("TOPLIVO_ACCESS_TOKEN", "", -1, "/", "localhost", config.App.Production, true)
			c.SetCookie("TOPLIVO_IDENTITY_TOKEN", "", -1, "/", "localhost", config.App.Production, true)
			c.Next()
			return
		}

		// Session is valid, set user info in context
		c.Set("session", session)
		c.Next()
	}
}
