package myWebsocket

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// GetUserId 从Gin上下文获取用户id
func GetUserId(c *gin.Context) int {
	// 获取当前用户id(从JWT解析)
	userData, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return 0
	}

	// 断言
	userMap, ok := userData.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
		return 0
	}

	// 提取user_id并转换为整数
	userIDFloat, ok := userMap["user_id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return 0
	}
	userID := int(userIDFloat)
	return userID
}

// SendReminderToClient WebSocket连接发送提醒消息
func (h *Handler) SendReminderToClient(userID int, message string) error {
	conn, ok := wsConnections[userID]
	if !ok {
		return fmt.Errorf("WebSocket连接未找到，用户ID：%d", userID)
	}
	return h.WriteMessage(conn, []byte(message))
}

// WriteMessage WebSocket写入消息
func (h *Handler) WriteMessage(conn *websocket.Conn, message []byte) error {
	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("发送WebSocket消息失败:", err)
	}
	return err
}
