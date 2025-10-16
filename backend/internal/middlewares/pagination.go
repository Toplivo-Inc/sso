// Package middlewares defines various middlewares
package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(defaultLimit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		limitQuery := c.Query("limit")
		limit, err := strconv.Atoi(limitQuery)
		if err != nil {
			limit = defaultLimit
		}

		pageQuery := c.Query("page")
		page, err := strconv.Atoi(pageQuery)
		if err != nil {
			page = 1
		}

		c.Set("limit", limit)
		c.Set("page", page)
		c.Next()
	}
}
