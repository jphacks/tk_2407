package middleware

import (
	"backend/pkg/jwt"
	"backend/pkg/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware middleware to verify JWT token from cookie
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Verify the token
		config := settings.Get()
		claims, err := jwt.Verify(tokenString, config.Service.Authentication.JwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Set the user information in the context
		c.Set("userID", claims.Id)
		c.Set("username", claims.Issuer)

		c.Next()
	}
}
