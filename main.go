package main

import (
	"backup/core"
	"backup/global"
	"go.uber.org/zap"
)

var (
	vip    = global.App.ConfigViper
	logger = global.App.ConfigZap
)

func main() {
	// 配置初始化
	vip = core.InitializeViper()
	// 日志初始化
	logger = core.InitializeZap()
	logger.Info("init test ", zap.String("zap_log", "zap_log"))
	core.RunServer()
}
