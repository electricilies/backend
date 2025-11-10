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
}

type RedisConfig struct {
	Enabled bool
	Image   string
}

type ContainersConfig struct {
	db       DBConfig
	keycloak KeycloakConfig
	minio    MinIOConfig
	redis    RedisConfig
}

func NewContainersConfig() *ContainersConfig {
	return &ContainersConfig{
		db: DBConfig{
			Enabled:  false,
			Image:    DBImage,
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
			Enabled:       false,
			Image:         KeycloakImage,
			AdminUsername: "electricilies",
			AdminPassword: "electricilies",
			ContextPath:   "/auth",
			RealmFilePath: filepath.Join("testdata", "electricilies-realm-export.json"),
		},
		minio: MinIOConfig{
			Enabled:  false,
			Image:    MinIOImage,
			User:     "electricilies",
			Password: "electricilies",
		},
		redis: RedisConfig{
			Enabled: false,
			Image:   RedisImage,
		},
	}
}

type Containers struct {
	Postgres *postgres.PostgresContainer
	Keycloak *keycloak.KeycloakContainer
	MinIO    *minio.MinioContainer
	Redis    *redis.RedisContainer
}

func NewContainers(ctx context.Context, cfg *ContainersConfig) (*Containers, error) {
	if cfg == nil {
		cfg = NewContainersConfig()
	}
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

func (c *Containers) Cleanup(t testing.TB) {
	for _, container := range []testcontainers.Container{
		c.Postgres,
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
