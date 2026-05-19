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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No auth header"})
		}
		token := strings.TrimPrefix(header, "Bearer ")
		userID, err := jw.GetID(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "failed to parse token"})
		}
		c.Set("userID", userID)
		c.Next()
	}
}
