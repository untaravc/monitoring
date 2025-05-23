package middleware

import (
	"monitoring/app/utils/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Only check Content-Type for PATCH and POST methods
		if c.Request.Method == http.MethodPatch || c.Request.Method == http.MethodPost {
			// If Content Type is empty or not application/json
			if c.ContentType() != "application/json" && c.ContentType() != "multipart/form-data" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported Media Type"})
				c.Abort()
				return
			}
		}

		// If There is no Bearer String
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := jwt.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		jwt_auth := jwt.ClaimData{Id: claims.Id, Email: claims.Email, Phone: claims.Phone}
		// You can set the claims to the context if needed
		c.Set("auth_user", jwt_auth)

		c.Next()
	}
}
