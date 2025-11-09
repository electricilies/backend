package integration

import (
	"context"
	"log"
	"path/filepath"

	keycloak "github.com/stillya/testcontainers-keycloak"
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
}

type RedisConfig struct {
	Enabled bool
	Image   string
}

type Config struct {
	db       DBConfig
	keycloak KeycloakConfig
	minio    MinIOConfig
	redis    RedisConfig
}

func NewConfig() *Config {
	return &Config{
		db: DBConfig{
			Enabled:  true,
			Image:    "paradedb/paradedb:17-v0.18.11",
			User:     "electricilies",
			Password: "electricilies",
			Database: "electricilies",
			InitScriptsPaths: []string{
				filepath.Join("testdata", "init.sql"),
				filepath.Join("testdata", "schema.sql"),
				filepath.Join("testdata", "trigger.sql"),
				filepath.Join("testdata", "seed.sql"),
				filepath.Join("testdata", "seed-fake.sql"),
				filepath.Join("testdata", "paradedb-index.sql"),
			},
		},
		keycloak: KeycloakConfig{
			Enabled:       true,
			Image:         "keycloak/keycloak:26.4.4",
			AdminUsername: "electricilies",
			AdminPassword: "electricilies",
			ContextPath:   "/auth",
			RealmFilePath: filepath.Join("testdata", "electricilies-realm-export.json"),
		},
		minio: MinIOConfig{
			Enabled:  true,
			Image:    "minio/minio:RELEASE.2025-09-07T16-13-09Z",
			User:     "electricilies",
			Password: "electricilies123",
		},
		redis: RedisConfig{
			Enabled: true,
			Image:   "redis:8.2.2-alpine3.22",
		},
	}
}

type Containers struct {
	Postgres *postgres.PostgresContainer
	Keycloak *keycloak.KeycloakContainer
	MinIO    *minio.MinioContainer
	Redis    *redis.RedisContainer
}

func NewContainers(ctx context.Context, cfg *Config) (*Containers, error) {
	var postgresContainer *postgres.PostgresContainer
	var keycloakContainer *keycloak.KeycloakContainer
	var minioContainer *minio.MinioContainer
	var redisContainer *redis.RedisContainer
	var err error

	if cfg.db.Enabled {
		postgresContainer, err = postgres.Run(
			ctx,
			cfg.db.Image,
			postgres.WithDatabase(cfg.db.Database),
			postgres.WithUsername(cfg.db.User),
			postgres.WithPassword(cfg.db.Password),
			postgres.WithOrderedInitScripts(cfg.db.InitScriptsPaths...),
			postgres.BasicWaitStrategies(),
		)
		if err != nil {
			log.Printf("failed to start postgres container: %v", err)
			return nil, err
		}
	}

	if cfg.keycloak.Enabled {
		keycloakContainer, err = keycloak.Run(
			ctx,
			cfg.keycloak.Image,
			keycloak.WithAdminUsername(cfg.keycloak.AdminUsername),
			keycloak.WithAdminPassword(cfg.keycloak.AdminPassword),
			keycloak.WithContextPath(cfg.keycloak.ContextPath),
			keycloak.WithRealmImportFile(cfg.keycloak.RealmFilePath),
		)
		if err != nil {
			log.Printf("failed to start keycloak container: %v", err)
			return nil, err
		}
	}

	if cfg.minio.Enabled {
		minioContainer, err = minio.Run(
			ctx,
			cfg.minio.Image,
			minio.WithUsername(cfg.minio.User),
			minio.WithPassword(cfg.minio.Password),
		)
		if err != nil {
			log.Printf("failed to start minio container: %v", err)
			return nil, err
		}
	}

	if cfg.redis.Enabled {
		redisContainer, err = redis.Run(
			ctx,
			cfg.redis.Image,
		)
		if err != nil {
			log.Printf("failed to start redis container: %v", err)
			return nil, err
		}
	}

	containers := &Containers{
		Postgres: postgresContainer,
		Keycloak: keycloakContainer,
		MinIO:    minioContainer,
		Redis:    redisContainer,
	}
	return containers, nil
}
