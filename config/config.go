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
	KcClientID     = "KC_CLIENT_ID"
	KcClientSecret = "KC_CLIENT_SECRET"
	KcRealm        = "KC_REALM"
	KcBasePath     = "KC_BASE_PATH"
	RedisAddr      = "REDIS_ADDRESS"
	S3AccessKey    = "S3_ACCESS_KEY"
	S3SecretKey    = "S3_SECRET_KEY"
	S3RegionName   = "S3_REGION_NAME"
	S3Endpoint     = "S3_ENDPOINT"
	S3Bucket       = "S3_BUCKET"
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
	KcClientID     string
	KcClientSecret string
	KcRealm        string
	KcBasePath     string
	RedisAddr      string
	S3AccessKey    string
	S3SecretKey    string
	S3RegionName   string
	S3Endpoint     string
	S3Bucket       string
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
		KcClientID:     viper.GetString(KcClientID),
		KcClientSecret: viper.GetString(KcClientSecret),
		KcRealm:        viper.GetString(KcRealm),
		KcBasePath:     viper.GetString(KcBasePath),
		RedisAddr:      viper.GetString(RedisAddr),
		S3AccessKey:    viper.GetString(S3AccessKey),
		S3SecretKey:    viper.GetString(S3SecretKey),
		S3RegionName:   viper.GetString(S3RegionName),
		S3Endpoint:     viper.GetString(S3Endpoint),
		S3Bucket:       viper.GetString(S3Bucket),
	}
}
