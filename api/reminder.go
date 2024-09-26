package api

import (
	"calendar-reminder/config"
	"calendar-reminder/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateReminder 新增提醒信息
func CreateReminder(c *gin.Context) {
	// 获取当前用户ID（从JWT解析）
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

	var request struct {
		Content        string `json:"content" binding:"required"`
		ReminderAt     string `json:"reminder_at" binding:"required"`
		ReminderMethod string `json:"reminder_method" binding:"required"`
	}

	// 绑定数据
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 新增提醒
	if err :=
		service.CreateReminder(config.DB, userID, request.Content, request.ReminderAt, request.ReminderMethod); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reminder"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reminder created successfully"})
}

// GetReminders 获取当前用户的提醒信息列表
func GetReminders(c *gin.Context) {
	// 获取当前用户ID（从JWT解析）
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

	// 获取提醒信息列表
	reminders, err := service.GetRemindersByUserID(config.DB, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reminders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reminders": reminders})
}

// UpdateReminder 更新提醒信息
func UpdateReminder(c *gin.Context) {
	// 获取当前用户ID（从JWT解析）
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

	// 获取提醒ID
	reminderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reminder ID"})
		return
	}

	var request struct {
		Content    string `json:"content" binding:"required"`
		ReminderAt string `json:"reminder_at" binding:"required"`
	}

	// 绑定数据
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 更新提醒
	err = service.UpdateReminder(config.DB, userID, reminderID, request.Content, request.ReminderAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reminder"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reminder updated successfully"})
}

// DeleteReminder 删除提醒信息
func DeleteReminder(c *gin.Context) {
	// 获取当前用户ID（从JWT解析）
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

	// 获取提醒ID
	reminderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reminder ID"})
		return
	}

	// 删除提醒
	err = service.DeleteReminder(config.DB, userID, reminderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete reminder"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reminder deleted successfully"})
}
