package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	configDBUsername          = "DB_USERNAME"
	configDBPassword          = "DB_PASSWORD"
	configDBHost              = "DB_HOST"
	configDBPort              = "DB_PORT"
	configDBName              = "DB_DATABASE"
	configDBURL               = "DB_URL"
	configEnvApp              = "ENV_APP"
	configLogStdout           = "LOG_ENABLE_STDOUT"
	configLogFile             = "LOG_ENABLE_FILE"
	configKCClientId          = "KC_CLIENT_ID"
	configKCClientSecret      = "KC_CLIENT_SECRET"
	configKCRealm             = "KC_REALM"
	configKCBasePath          = "KC_BASE_PATH"
	configKCHttpManagmentPath = "KC_HTTP_MANAGEMENT_PATH"
	configRedisAddr           = "REDIS_ADDRESS"
	configS3AccessKey         = "S3_ACCESS_KEY"
	configS3SecretKey         = "S3_SECRET_KEY"
	configS3RegionName        = "S3_REGION_NAME"
	configS3Endpoint          = "S3_ENDPOINT"
	configS3Bucket            = "S3_BUCKET"
	configTimeZone            = "TIMEZONE"
	configSwaggerEnv          = "SWAGGER_ENV"
	configPublicKeycloakURL   = "PUBLIC_KEYCLOAK_URL"
)

type Config struct {
	DBUsername          string
	DBPassword          string
	DBHost              string
	DBPort              int
	DBName              string
	DBURL               string
	EnvApp              string
	EnableStdout        bool
	EnableFile          bool
	KCClientId          string
	KCClientSecret      string
	KCRealm             string
	KCBasePath          string
	KCHttpManagmentPath string
	RedisAddr           string
	S3AccessKey         string
	S3SecretKey         string
	S3RegionName        string
	S3Endpoint          string
	S3Bucket            string
	TimeZone            string
	SwaggerEnv          string
	PublicKeycloakURL   string
}

func ProvideConfig() *Config {
	viper.AutomaticEnv()

	viper.SetDefault(configDBPort, 5432)
	viper.SetDefault(configLogStdout, true)
	viper.SetDefault(configLogFile, false)

	viper.SetDefault(configTimeZone, "Asia/Ho_Chi_Minh")
	if viper.GetString(configS3Bucket) == "" {
		log.Print("You need to set S3_BUCKET environment variable")
	}

	return &Server{
		DBUsername:          viper.GetString(configDBUsername),
		DBPassword:          viper.GetString(configDBPassword),
		DBHost:              viper.GetString(configDBHost),
		DBPort:              viper.GetInt(configDBPort),
		DBName:              viper.GetString(configDBName),
		DBURL:               viper.GetString(configDBURL),
		EnvApp:              viper.GetString(configEnvApp),
		EnableStdout:        viper.GetBool(configLogStdout),
		EnableFile:          viper.GetBool(configLogFile),
		KCClientId:          viper.GetString(configKCClientId),
		KCClientSecret:      viper.GetString(configKCClientSecret),
		KCRealm:             viper.GetString(configKCRealm),
		KCBasePath:          viper.GetString(configKCBasePath),
		KCHttpManagmentPath: viper.GetString(configKCHttpManagmentPath),
		RedisAddr:           viper.GetString(configRedisAddr),
		S3AccessKey:         viper.GetString(configS3AccessKey),
		S3SecretKey:         viper.GetString(configS3SecretKey),
		S3RegionName:        viper.GetString(configS3RegionName),
		S3Endpoint:          viper.GetString(configS3Endpoint),
		S3Bucket:            viper.GetString(configS3Bucket),
		TimeZone:            viper.GetString(configTimeZone),
		SwaggerEnv:          viper.GetString(configSwaggerEnv),
		PublicKeycloakURL:   viper.GetString(configPublicKeycloakURL),
	}
}
