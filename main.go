package main

import (
	"short-url-go/internal/global"
	"short-url-go/internal/initialize"
)

func main() {
	global.G_VP = initialize.Viper() // 初始化Viper
	global.G_LOG = initialize.Zap()  // 初始化zap日志库
	global.G_DB = initialize.Gorm()  // gorm连接数据库
	initialize.Routers()             // 初始化路由
	initialize.RunWindowsServer()    // 运行服务
}
