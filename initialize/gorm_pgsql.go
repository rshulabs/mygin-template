package initialize

import (
	"backup/global"
	"backup/initialize/internal"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormPgsql() *gorm.DB {
	p := global.App.Config.PGSQL
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(),
		PreferSimpleProtocol: false,
	}
	db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular))
	if err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		fmt.Println("gorm link postgresql success")
		return db
	}
}
