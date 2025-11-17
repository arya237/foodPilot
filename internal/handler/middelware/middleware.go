package middelware

import (
	"log"
	"net/http"
	"strings"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		claims, err := ValidateJWT(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("token", claims.Token)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func LimitMiddelware(limit *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		context, err := limit.Get(c, clientIP)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if context.Reached {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":     "Too many requests, please try again later.",
				"remaining": context.Remaining,
			})
		}

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}

		role, ok := roleVal.(models.UserRole)
		log.Println(role, ok)
		if ok && role == models.RoleAdmin {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "admin only"})
		c.Abort()
	}
}
