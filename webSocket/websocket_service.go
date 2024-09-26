package websocket

import (
	"fmt"
)

// SendReminderToClient 方法用于向特定用户的WebSocket连接发送提醒消息
func (h *WebSocketHandler) SendReminderToClient(userID int, message string) error {
	conn, ok := wsConnections[userID]
	if !ok {
		return fmt.Errorf("WebSocket连接未找到，用户ID：%d", userID)
	}
	return h.WriteMessage(conn, []byte(message))
}
