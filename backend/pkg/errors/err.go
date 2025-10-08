package errors

import (
	"log"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (a *AppError) Error() string {
	if a.Err != nil {
		return a.Err.Error()
	}
	return a.Message
}

func AppErr(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			log.Println("ERROR:", err.Error())

			appErr, ok := err.(*AppError)
			if !ok {
				msg := err.Error()
				if msg == "" {
					msg = "internal server error"
				}
				c.AbortWithStatusJSON(500, gin.H{
					"message": msg,
				})
				return
			}
			c.AbortWithStatusJSON(appErr.Code, gin.H{
				"message": appErr.Message,
			})
		}
	}
}

