package config

import (
	"github.com/spf13/viper"
)

const (
	KeyDBUsername = "DB_USERNAME"
	KeyDBPassword = "DB_PASSWORD"
	KeyDBHost     = "DB_HOST"
	KeyDBPort     = "DB_PORT"
	KeyDBName     = "DB_DATABASE"
	KeyEnvApp     = "ENV_APP"
	KeyLogStdout  = "LOG_ENABLE_STDOUT"
	KeyLogFile    = "LOG_ENABLE_FILE"
)

type Config struct {
	DBUsername   string
	DBPassword   string
	DBHost       string
	DBPort       int
	DBName       string
	EnvApp       string
	EnableStdout bool
	EnableFile   bool
}

var Cfg *Config

func LoadConfig() {
	viper.AutomaticEnv()

	viper.SetDefault(KeyDBPort, 5432)
	viper.SetDefault(KeyLogStdout, true)
	viper.SetDefault(KeyLogFile, false)

	Cfg = &Config{
		DBUsername:   viper.GetString(KeyDBUsername),
		DBPassword:   viper.GetString(KeyDBPassword),
		DBHost:       viper.GetString(KeyDBHost),
		DBPort:       viper.GetInt(KeyDBPort),
		DBName:       viper.GetString(KeyDBName),
		EnvApp:       viper.GetString(KeyEnvApp),
		EnableStdout: viper.GetBool(KeyLogStdout),
		EnableFile:   viper.GetBool(KeyLogFile),
	}
}
