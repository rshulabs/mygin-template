package initialize

import (
	"backup/global"
	"backup/initialize/internal"
	"backup/model"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.App.Config.MySQL
	if m.Dbname == "" {
		return nil
	}

	// 创建 mysql.Config 实例，其中包含了连接数据库所需的信息，比如 DSN (数据源名称)，字符串类型字段的默认长度以及自动根据版本进行初始化等参数。
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	// 打开数据库连接 驱动 gorm配置
	db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular))

	// 将引擎设置为我们配置的引擎，并设置每个连接的最大空闲数和最大连接数。
	if err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)

		fmt.Println(" gorm link mysql success")
		initMySqlTables(db)
		return db
	}
}

// 数据库初始化
func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		global.App.ConfigZap.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}
