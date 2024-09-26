package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ParseToken 中间件用于验证 JWT Token
func ParseToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 解析 Token
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		jwtService := &JWT{}
		userID, err := jwtService.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 将用户 ID 存储到上下文中，以便后续使用
		c.Set("userID", userID)
		c.Next()
	}
}
