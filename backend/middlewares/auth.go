package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rc.justin.ooo/models"
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
			data, _ := json.Marshal(user)
			fmt.Println(string(data))
			c.Next()
			return
		}

		c.AbortWithStatus(http.StatusUnauthorized)

	}
}
