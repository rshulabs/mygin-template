package initialize

/*
1.先定义数据库配置，和全局
2.gorm配置，日志级别，输出格式，写入器
3.分别定义不同数据库初始化方法
4.按需初始化gorm，绑定全局
*/
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
