package global

import (
	"sdwan.ink/sdwan/config"
	"gorm.io/gorm"
    "github.com/spf13/viper"
    "go.uber.org/zap"
)


var (
    WanConfig      config.Config
    WanViper       *viper.Viper
    WanLogger      *zap.Logger
    WanDb          *gorm.DB
)
