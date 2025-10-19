package config

import (
	"github.com/spf13/viper"
)

const (
	DBUsername     = "DB_USERNAME"
	DBPassword     = "DB_PASSWORD"
	DBHost         = "DB_HOST"
	DBPort         = "DB_PORT"
	DBName         = "DB_DATABASE"
	EnvApp         = "ENV_APP"
	LogStdout      = "LOG_ENABLE_STDOUT"
	LogFile        = "LOG_ENABLE_FILE"
	KCClientID     = "KC_CLIENT_ID"
	KCClientSecret = "KC_CLIENT_SECRET"
	KCRealm        = "KC_REALM"
)

type Config struct {
	DBUsername     string
	DBPassword     string
	DBHost         string
	DBPort         int
	DBName         string
	EnvApp         string
	EnableStdout   bool
	EnableFile     bool
	KCClientID     string
	KCClientSecret string
	KCRealm        string
}

var Cfg *Config

func LoadConfig() {
	viper.AutomaticEnv()

	viper.SetDefault(DBPort, 5432)
	viper.SetDefault(LogStdout, true)
	viper.SetDefault(LogFile, false)

	Cfg = &Config{
		DBUsername:     viper.GetString(DBUsername),
		DBPassword:     viper.GetString(DBPassword),
		DBHost:         viper.GetString(DBHost),
		DBPort:         viper.GetInt(DBPort),
		DBName:         viper.GetString(DBName),
		EnvApp:         viper.GetString(EnvApp),
		EnableStdout:   viper.GetBool(LogStdout),
		EnableFile:     viper.GetBool(LogFile),
		KCClientID:     viper.GetString(KCClientID),
		KCClientSecret: viper.GetString(KCClientSecret),
		KCRealm:        viper.GetString(KCRealm),
	}
}
