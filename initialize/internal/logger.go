package internal

import (
	"backup/global"
	"fmt"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.App.Config.App.DbType {
	case "mysql":
		logZap = global.App.Config.MySQL.LogZap
	}
	if logZap {
		global.App.ConfigZap.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
