//go:build integration

package component

import (
	"context"
	"log"
	"path/filepath"
	"sort"
	"testing"

	keycloak "github.com/stillya/testcontainers-keycloak"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/minio"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

type DBConfig struct {
	Image            string
	User             string
	Password         string
	Database         string
	InitScriptsPaths []string
	Seed             bool
	SeedScriptsPaths []string
}

type KeycloakConfig struct {
	Image         string
	AdminUsername string
	AdminPassword string
	ContextPath   string
	RealmFilePath string
}

type MinIOConfig struct {
	Image    string
	User     string
	Password string
	Bucket   string
}

type RedisConfig struct {
	Image string
}

type ContainersConfig struct {
	DB       *DBConfig
	Keycloak *KeycloakConfig
	MinIO    *MinIOConfig
	Redis    *RedisConfig
}

type NewContainersConfigParam struct {
	DBEnabled       bool
	KeycloakEnabled bool
	MinIOEnabled    bool
	RedisEnabled    bool
}

func NewContainersConfig(param *NewContainersConfigParam) *ContainersConfig {
	var db *DBConfig
	var keycloak *KeycloakConfig
	var minIO *MinIOConfig
	var redis *RedisConfig
	if param == nil {
		param = &NewContainersConfigParam{}
	}
	if param.DBEnabled {
		db = &DBConfig{
			Image:    DBImage,
			User:     "electricilies",
			Password: "electricilies",
			Database: "electricilies",
			InitScriptsPaths: []string{
				filepath.Join("..", "testdata", "00-init.sql"),
				filepath.Join("..", "testdata", "20-schema.sql"),
				filepath.Join("..", "testdata", "21-trigger.sql"),
				filepath.Join("..", "testdata", "40-paradedb-index.sql"),
			},
			Seed: false,
			SeedScriptsPaths: []string{
				filepath.Join("..", "testdata", "30-seed.sql"),
			},
		}
	}
	if param.KeycloakEnabled {
		keycloak = &KeycloakConfig{
			Image:         KeycloakImage,
			AdminUsername: "electricilies",
			AdminPassword: "electricilies",
			ContextPath:   "/auth",
			RealmFilePath: filepath.Join("..", "testdata", "electricilies-realm-export.json"),
		}
	}
	if param.MinIOEnabled {
		minIO = &MinIOConfig{
			Image:    MinIOImage,
			User:     "electricilies",
			Password: "electricilies",
			Bucket:   "electricilies",
		}
	}
	if param.RedisEnabled {
		redis = &RedisConfig{
			Image: RedisImage,
		}
	}
	return &ContainersConfig{
		DB:       db,
		Keycloak: keycloak,
		MinIO:    minIO,
		Redis:    redis,
	}
}

type Containers struct {
	DB       *postgres.PostgresContainer
	Keycloak *keycloak.KeycloakContainer
	MinIO    *minio.MinioContainer
	Redis    *redis.RedisContainer
}

func setupPostgres(ctx context.Context, cfg *DBConfig) (*postgres.PostgresContainer, error) {
	if cfg.Seed {
		cfg.InitScriptsPaths = append(cfg.InitScriptsPaths, cfg.SeedScriptsPaths...)
		sort.Strings(cfg.InitScriptsPaths)
	}

	postgresContainer, err := postgres.Run(
		ctx,
		cfg.Image,
		postgres.WithDatabase(cfg.Database),
		postgres.WithUsername(cfg.User),
		postgres.WithPassword(cfg.Password),
		postgres.WithInitScripts(cfg.InitScriptsPaths...),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		log.Printf("failed to start postgres container: %v", err)
		return nil, err
	}
	return postgresContainer, nil
}

func setupKeycloak(ctx context.Context, cfg *KeycloakConfig) (*keycloak.KeycloakContainer, error) {
	keycloakContainer, err := keycloak.Run(
		ctx,
		cfg.Image,
		keycloak.WithAdminUsername(cfg.AdminUsername),
		keycloak.WithAdminPassword(cfg.AdminPassword),
		keycloak.WithContextPath(cfg.ContextPath),
		keycloak.WithRealmImportFile(cfg.RealmFilePath),
	)
	if err != nil {
		log.Printf("failed to start keycloak container: %v", err)
		return nil, err
	}
	return keycloakContainer, nil
}

func setupMinio(ctx context.Context, cfg *MinIOConfig) (*minio.MinioContainer, error) {
	minioContainer, err := minio.Run(
		ctx,
		cfg.Image,
		minio.WithUsername(cfg.User),
		minio.WithPassword(cfg.Password),
	)
	if err != nil {
		log.Printf("failed to start minio container: %v", err)
		return nil, err
	}
	return minioContainer, nil
}

func setupRedis(ctx context.Context, cfg *RedisConfig) (*redis.RedisContainer, error) {
	redisContainer, err := redis.Run(
		ctx,
		cfg.Image,
	)
	if err != nil {
		log.Printf("failed to start redis container: %v", err)
		return nil, err
	}
	return redisContainer, nil
}

func NewContainers(ctx context.Context, cfg *ContainersConfig) (*Containers, error) {
	if cfg == nil {
		cfg = NewContainersConfig(nil)
	}
	var postgresContainer *postgres.PostgresContainer
	var keycloakContainer *keycloak.KeycloakContainer
	var minioContainer *minio.MinioContainer
	var redisContainer *redis.RedisContainer
	var err error

	if cfg.DB != nil {
		postgresContainer, err = setupPostgres(ctx, cfg.DB)
		if err != nil {
			return nil, err
		}
	}

	if cfg.Keycloak != nil {
		keycloakContainer, err = setupKeycloak(ctx, cfg.Keycloak)
		if err != nil {
			return nil, err
		}
	}

	if cfg.MinIO != nil {
		minioContainer, err = setupMinio(ctx, cfg.MinIO)
		if err != nil {
			return nil, err
		}
	}

	if cfg.Redis != nil {
		redisContainer, err = setupRedis(ctx, cfg.Redis)
		if err != nil {
			return nil, err
		}
	}

	containers := &Containers{
		DB:       postgresContainer,
		Keycloak: keycloakContainer,
		MinIO:    minioContainer,
		Redis:    redisContainer,
	}
	return containers, nil
}

func (c *Containers) Cleanup(t testing.TB) {
	for _, container := range []testcontainers.Container{
		c.DB,
		c.Keycloak,
		c.MinIO,
		c.Redis,
	} {
		if container == nil {
			continue
		}
		testcontainers.CleanupContainer(t, container)
	}
}
