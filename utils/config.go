package utils

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type ConfigInst struct {
	PORT            int           `mapstructure:"PORT"`
	DatabaseURL     string        `mapstructure:"DATABASE_URL"`
	SessionDuration time.Duration `mapstructure:"SESSION_DURATION"`
	CookieEncryptionKey string `mapstructure:"COOKIE_ENCRYPTION_KEY"`
	SessionCookieName string `mapstructure:"COOKIE_SESSION_NAME"`
}

var Config ConfigInst

func InitConfig() *ConfigInst {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &Config
}
