package middleware

import (
	"ecommerce/user-service/internal/jw"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetID() gin.HandlerFunc {
	return func(c *gin.Context) {
		jw := jw.NewJWT()
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
		userID, err := jw.GetID(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token."})
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
