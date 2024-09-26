package main

import (
	"calendar-reminder/config"
	"calendar-reminder/router"
	"calendar-reminder/service"
	"fmt"
)

func main() {
	// 初始化数据库
	config.InitDB()
	config.InitRedis()

	// 开启定时器
	service.Init()

	// 初始化路由
	r := router.SetupRouter()

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		fmt.Println("启动失败")
	}
}
