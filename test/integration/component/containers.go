//go:build integration

package component

import (
	"context"
	"log"
	"path/filepath"
	"testing"

	keycloak "github.com/stillya/testcontainers-keycloak"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/minio"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

type DBConfig struct {
	Enabled          bool
	Image            string
	User             string
	Password         string
	Database         string
	InitScriptsPaths []string
}

type KeycloakConfig struct {
	Enabled       bool
	Image         string
	AdminUsername string
	AdminPassword string
	ContextPath   string
	RealmFilePath string
}

type MinIOConfig struct {
	Enabled  bool
	Image    string
	User     string
	Password string
	Bucket   string
}

type RedisConfig struct {
	Enabled bool
	Image   string
}

type ContainersConfig struct {
	DB       DBConfig
	Keycloak KeycloakConfig
	MinIO    MinIOConfig
	Redis    RedisConfig
}

func NewContainersConfig() *ContainersConfig {
	return &ContainersConfig{
		DB: DBConfig{
			Enabled:  false,
			Image:    DBImage,
			User:     "electricilies",
			Password: "electricilies",
			Database: "electricilies",
			InitScriptsPaths: []string{
				filepath.Join("..", "testdata", "00-init.sql"),
				filepath.Join("..", "testdata", "20-schema.sql"),
				filepath.Join("..", "testdata", "21-trigger.sql"),
				filepath.Join("..", "testdata", "30-seed.sql"),
				filepath.Join("..", "testdata", "40-paradedb-index.sql"),
			},
		},
		Keycloak: KeycloakConfig{
			Enabled:       false,
			Image:         KeycloakImage,
			AdminUsername: "electricilies",
			AdminPassword: "electricilies",
			ContextPath:   "/auth",
			RealmFilePath: filepath.Join("..", "testdata", "electricilies-realm-export.json"),
		},
		MinIO: MinIOConfig{
			Enabled:  false,
			Image:    MinIOImage,
			User:     "electricilies",
			Password: "electricilies",
			Bucket:   "electricilies",
		},
		Redis: RedisConfig{
			Enabled: false,
			Image:   RedisImage,
		},
	}
}

type Containers struct {
	DB       *postgres.PostgresContainer
	Keycloak *keycloak.KeycloakContainer
	MinIO    *minio.MinioContainer
	Redis    *redis.RedisContainer
}

func setupPostgres(ctx context.Context, cfg DBConfig) (*postgres.PostgresContainer, error) {
	if !cfg.Enabled {
		return nil, nil
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

func setupKeycloak(ctx context.Context, cfg KeycloakConfig) (*keycloak.KeycloakContainer, error) {
	if !cfg.Enabled {
		return nil, nil
	}

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

func setupMinio(ctx context.Context, cfg MinIOConfig) (*minio.MinioContainer, error) {
	if !cfg.Enabled {
		return nil, nil
	}

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

func setupRedis(ctx context.Context, cfg RedisConfig) (*redis.RedisContainer, error) {
	if !cfg.Enabled {
		return nil, nil
	}

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
		cfg = NewContainersConfig()
	}

	postgresContainer, err := setupPostgres(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}

	keycloakContainer, err := setupKeycloak(ctx, cfg.Keycloak)
	if err != nil {
		return nil, err
	}

	minioContainer, err := setupMinio(ctx, cfg.MinIO)
	if err != nil {
		return nil, err
	}

	redisContainer, err := setupRedis(ctx, cfg.Redis)
	if err != nil {
		return nil, err
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
