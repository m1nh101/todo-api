package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m1ngi/todo-api/helper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Request.Cookie("access_token")
		if token == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		currentUserId, err := helper.ValidateJwtToken(token.Value)

		if err != nil {
			log.Fatal(err)
		}

		c.Set("Authorized", true)
		c.Set("UserId", currentUserId)
		c.Next()
	}
}
