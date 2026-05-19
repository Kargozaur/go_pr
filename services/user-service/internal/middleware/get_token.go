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
		c.Set("token", token)
		c.Next()
	}
}
