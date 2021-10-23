package middlewares

import (
	"net/http"

	"clientellapp.com/models"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Token")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := models.UserFromToken(token)
		if err == nil {
			c.Set("user", user)
			c.Next()
			return
		}

		c.AbortWithStatus(http.StatusUnauthorized)

	}
}
