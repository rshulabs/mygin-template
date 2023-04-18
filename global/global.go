package global

import (
	"backup/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config      config.Configuration
	ConfigViper *viper.Viper
	ConfigZap   *zap.Logger
	ConfigDB    *gorm.DB
	Redis       *redis.Client
}

var App = new(Application)
