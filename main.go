package main

import (
	"backup/core"
	"backup/global"
	"backup/initialize"
	"go.uber.org/zap"
)

var (
	vip    = global.App.ConfigViper
	logger = global.App.ConfigZap
	db     = global.App.ConfigDB
)

func main() {
	// 配置初始化
	vip = core.InitializeViper()
	// 日志初始化
	logger = core.InitializeZap()
	logger.Info("init start ", zap.String("zap_log", "zap_log"))
	// 数据库初始化
	db = initialize.Gorm()
	// 其他配置初始化
	initialize.OtherInit()
	// 启动服务
	core.RunServer()
}
