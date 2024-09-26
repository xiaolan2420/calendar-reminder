package myWebsocket

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 全局的 WebSocket 连接映射，使用用户id作为键
var wsConnections = make(map[int]*websocket.Conn)

// Handler 结构体用于处理 WebSocket 相关的操作
type Handler struct {
}

// HandleWebSocket 函数用于处理WebSocket连接相关逻辑
func (h *Handler) HandleWebSocket(c *gin.Context) {
	// 获取当前用户id（从JWT解析）
	userData, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// 断言
	userMap, ok := userData.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
		return
	}

	// 提取 user_id 并转换为整数
	userIDFloat, ok := userMap["user_id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	userID := int(userIDFloat)

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket 升级失败:", err)
		return
	}
	fmt.Println(conn)
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭 WebSocket 连接失败:", err)
		}
	}(conn)

	// 将连接存储到全局连接映射中
	wsConnections[userID] = conn

	fmt.Println("WebSocket 连接成功建立，用户 ID：", userID)

	// 处理 WebSocket 连接
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("读取 WebSocket 消息失败:", err)
			break
		}
		fmt.Println("接收到的 WebSocket 消息:", string(message))

		// 回复收到消息
		err = conn.WriteMessage(websocket.TextMessage, []byte("服务器收到消息"))
		if err != nil {
			fmt.Println("发送 WebSocket 消息失败:", err)
			break
		}
	}
}
