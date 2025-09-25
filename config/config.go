package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	EnvApp     string
	AppPort    string
}

var Cfg Config

func LoadConfig() {
	viper.AutomaticEnv()
	Cfg = Config{
		DBUsername: viper.GetString("DB_USERNAME"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetInt("DB_PORT"),
		DBName:     viper.GetString("DB_DATABASE"),
		EnvApp:     viper.GetString("ENV_APP"),
		AppPort:    viper.GetString("PORT"),
	}
}
