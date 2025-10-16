package middlewares

import "github.com/gin-gonic/gin"

type PathFlag byte

const (
	ID       PathFlag = iota
	SecondID PathFlag = iota
)

func Path(flags PathFlag) gin.HandlerFunc {
	return func(c *gin.Context) {
		if (flags & ID) != 0 {
			id := c.Param("id")
			c.Set("id", id)
		}

		if (flags & SecondID) != 0 {
			id := c.Param("second_id")
			c.Set("second_id", id)
		}

		c.Next()
	}
}
