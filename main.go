package main

import (
	"backup/core"
	"backup/global"
	"backup/initialize"
	"go.uber.org/zap"
)

func main() {
	// 配置初始化
	global.App.ConfigViper = core.InitializeViper()
	// 日志初始化
	global.App.ConfigZap = core.InitializeZap()
	global.App.ConfigZap.Info("init start ", zap.String("zap_log", "zap_log"))
	// 数据库初始化
	global.App.ConfigDB = initialize.Gorm()
	// 其他配置初始化
	initialize.OtherInit()
	// 启动服务
	core.RunServer()
}
