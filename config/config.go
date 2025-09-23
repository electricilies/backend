package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl string
}

var Cfg Config

func LoadConfig() {
	viper.AutomaticEnv()
	Cfg = Config{
		DBUrl: viper.GetString("DB_URL"),
	}
}
