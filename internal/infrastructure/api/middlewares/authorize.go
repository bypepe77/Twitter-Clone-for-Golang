package middlewares

import (
	"net/http"
	"strings"

	jwtManager "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	BearerSchema = "Bearer "
)

func Authorize(manager jwtManager.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) <= len(BearerSchema) || !strings.HasPrefix(authHeader, BearerSchema) {
			unauthorized(c)
			return
		}

		tokenString := authHeader[len(BearerSchema):]
		token, err := manager.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			unauthorized(c)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			unauthorized(c)
			return
		}

		finalClaims, err := manager.FromMapClaims(claims)
		if err != nil || finalClaims == nil {
			unauthorized(c)
			return
		}

		c.Set("claims", finalClaims)
	}
}

func unauthorized(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}
