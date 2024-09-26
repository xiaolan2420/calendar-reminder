package api

import (
	"calendar-reminder/config"
	"calendar-reminder/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCode 获取验证码
func GetCode(c *gin.Context) {
	var request struct {
		Phone string `json:"phone" binding:"required"`
	}

	// 绑定数据
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := service.GetCode(request.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "send code successfully"})

}

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	var request struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
		Code     string `json:"code" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	// 绑定数据
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 注册用户
	err := service.RegisterUser(config.DB, request.Phone, request.Password, request.Email, request.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginUserByPassword(c *gin.Context) {
	var request struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// 绑定数据
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 登录
	token, err := service.LoginUser(config.DB, request.Phone, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 返回 JWT Token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// LoginUserByCode 使用验证码登录
func LoginUserByCode(c *gin.Context) {
	var request struct {
		Phone string `json:"phone" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}

	// 绑定数据
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 登录
	token, err := service.LoginUserByCode(config.DB, request.Phone, request.Code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 返回 JWT Token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
