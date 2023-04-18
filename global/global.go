package global

import (
	"backup/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config      config.Configuration
	ConfigViper *viper.Viper
	ConfigZap   *zap.Logger
	ConfigDB    *gorm.DB
}

var App = new(Application)
