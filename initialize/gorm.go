package initialize

import (
	"backup/global"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.App.Config.App.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgsql()
	default:
		return GormMysql()
	}
}
