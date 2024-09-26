package router

import (
	"calendar-reminder/api"
	"calendar-reminder/myWebsocket"
	"calendar-reminder/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 用户注册和登录
	r.GET("/getCode", api.GetCode)
	r.POST("/register", api.RegisterUser)
	r.POST("/loginByPassword", api.LoginUserByPassword)
	r.POST("/loginByCode", api.LoginUserByCode)

	handler := myWebsocket.Handler{}
	r.GET("/ws", handler.HandleWebSocket)

	// 提醒管理路由
	authorized := r.Group("/")
	authorized.Use(utils.ParseToken()) // 使用 JWT 中间件
	{
		authorized.POST("/reminders", api.CreateReminder)
		authorized.GET("/reminders", api.GetReminders)
		authorized.PUT("/reminders/:id", api.UpdateReminder)
		authorized.DELETE("/reminders/:id", api.DeleteReminder)

	}

	return r
}
