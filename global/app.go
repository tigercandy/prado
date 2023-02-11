package global

import (
	"github.com/spf13/viper"
	"github.com/tigercandy/prado/configs"
	"go.uber.org/zap"
)

const Version = "1.0.0"

type Application struct {
	ConfigViper *viper.Viper
	Config      configs.Configuration
	Log         *zap.Logger
}

var App = new(Application)
