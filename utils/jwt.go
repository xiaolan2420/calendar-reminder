package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	// 签名密钥
	secretKey = "zhu"
)

// JWT 结构体
type JWT struct {
	// 没有字段，所有操作通过方法完成
}

// NewJWT 创建 JWT 实例
func NewJWT() *JWT {
	return &JWT{}
}

// GenerateToken 生成 token 用户id和token过期时间
func (j *JWT) GenerateToken(userID int) (string, error) {

	expiresAt := time.Now().Add(time.Hour * 1)

	// 创建 JWT 的声明
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expiresAt.Unix(),
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseToken 解析和验证 token
func (j *JWT) ParseToken(tokenStr string) (jwt.MapClaims, error) {
	// 解析 token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法与我们预期的一致
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 提取声明
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
