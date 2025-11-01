package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	DbUsername     = "DB_USERNAME"
	DbPassword     = "DB_PASSWORD"
	DbHost         = "DB_HOST"
	DbPort         = "DB_PORT"
	DbName         = "DB_DATABASE"
	EnvApp         = "ENV_APP"
	LogStdout      = "LOG_ENABLE_STDOUT"
	LogFile        = "LOG_ENABLE_FILE"
	KcClientId     = "KC_CLIENT_ID"
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
	DbUsername     string
	DbPassword     string
	DbHost         string
	DbPort         int
	DbName         string
	EnvApp         string
	EnableStdout   bool
	EnableFile     bool
	KcClientId     string
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
	if viper.GetString(S3Bucket) == "" {
		log.Print("You need to set S3_BUCKET environment variable")
	}
	viper.AutomaticEnv()

	viper.SetDefault(DbPort, 5432)
	viper.SetDefault(LogStdout, true)
	viper.SetDefault(LogFile, false)

	Cfg = &Config{
		DbUsername:     viper.GetString(DbUsername),
		DbPassword:     viper.GetString(DbPassword),
		DbHost:         viper.GetString(DbHost),
		DbPort:         viper.GetInt(DbPort),
		DbName:         viper.GetString(DbName),
		EnvApp:         viper.GetString(EnvApp),
		EnableStdout:   viper.GetBool(LogStdout),
		EnableFile:     viper.GetBool(LogFile),
		KcClientId:     viper.GetString(KcClientId),
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
