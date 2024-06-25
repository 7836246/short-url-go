package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"short-url-go/internal/config"
)

var (
	G_DB     *gorm.DB
	G_LOG    *zap.Logger
	G_CONFIG config.Server
	G_VP     *viper.Viper
)
