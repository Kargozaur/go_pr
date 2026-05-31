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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No auth header."})
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		fromCookie, err := c.Request.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to read a cookie."})
			return
		}
		if token != fromCookie.Value || fromCookie.Value == "" || fromCookie == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cookie doesn't match header token."})
			return
		}
		c.Set("token", token)
		c.Next()
	}
}
