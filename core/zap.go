package core

/*
思路
1.先完成zap配置，定义全局变量
2.创建log目录（判断是否存在）
3.日志编码器配置，日志写入器配置
4.初始化日志，绑定到全局zap
*/
import (
	"backup/core/internal"
	"backup/global"
	"backup/utils"
	"fmt"

	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitializeZap Zap 获取 zap.Logger
func InitializeZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.App.Config.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.App.Config.Zap.Director)
		_ = os.Mkdir(global.App.Config.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.App.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	fmt.Println("zap log init success")
	return logger
}
