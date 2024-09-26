package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// 定义全局的WebSocket升级器
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 全局的WebSocket连接映射，使用用户ID作为键
var wsConnections = make(map[int]*websocket.Conn)

// Handler 结构体用于处理WebSocket相关的操作
type Handler struct {
}

// HandleWebSocket 函数用于处理WebSocket连接相关逻辑
func HandleWebSocket() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从 Gin 上下文中获取用户 ID
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = r
		userID := GetUserId(ctx)
		if userID == 0 {
			fmt.Println("获取用户ID失败")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		conn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("WebSocket 升级失败:", err)
			return
		}
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				fmt.Println("关闭 WebSocket 连接失败:", err)
			}
		}(conn)

		// 将连接存储到全局连接映射中
		wsConnections[userID] = conn

		// 处理 WebSocket 连接
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("读取 WebSocket 消息失败:", err)
				break
			}
			fmt.Println("接收到的 WebSocket 消息:", string(message))

			// 这里可以添加发送消息的逻辑
			err = conn.WriteMessage(websocket.TextMessage, []byte("服务器收到消息"))
			if err != nil {
				fmt.Println("发送 WebSocket 消息失败:", err)
				break
			}
		}
	}
}
