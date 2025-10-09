package middlewares

import (
	"sso/internal/service"
	"sso/pkg/config"

	"github.com/gin-gonic/gin"
)

// SessionMiddleware checks whether TOPLIVO_SESSION_TOKEN is provided and valid
func SessionMiddleware(sessionService service.SessionService, config *config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        sessionToken, err := c.Cookie("TOPLIVO_SESSION_TOKEN")
        if err != nil {
            // No session cookie, user is not authenticated
            c.Next()
            return
        }

        session, err := sessionService.ValidateSession(sessionToken)
        if err != nil {
            // Invalid session, clear the cookie
            c.SetCookie("TOPLIVO_SESSION_TOKEN", "", -1, "/", "localhost", config.App.Production, true)
            c.Next()
            return
        }

        // Session is valid, set user info in context
        c.Set("session", session)
        c.Next()
    }
}
