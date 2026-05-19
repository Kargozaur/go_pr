package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No auth header"})
		}
		token := strings.TrimPrefix(header, "Bearer ")
		fromCookie, _ := c.Request.Cookie("access_token")
		if token != fromCookie.Value {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cookie doesn't match header token"})
		}
		c.Set("token", token)
		c.Next()
	}
}
