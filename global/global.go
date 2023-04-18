package global

import (
	"backup/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	Config      config.Configuration
	ConfigViper *viper.Viper
	ConfigZap   *zap.Logger
}

var App = new(Application)
