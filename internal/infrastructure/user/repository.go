package user

import (
	"backend/config"
	"backend/internal/constant"
	"backend/internal/domain/user"
	"backend/internal/infrastructure/errors"
	"backend/internal/infrastructure/presistence/postgres"
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redis/go-redis/v9"
)

type repositoryImpl struct {
	db             *postgres.Queries
	s3Client       *s3.Client
	redisClient    *redis.Client
	keycloakClient *gocloak.GoCloak
}

func NewRepository(query *postgres.Queries, s3Client *s3.Client, redisClient *redis.Client, keycloakClient *gocloak.GoCloak) user.Repository {
	return &repositoryImpl{
		db:             query,
		s3Client:       s3Client,
		redisClient:    redisClient,
		keycloakClient: keycloakClient,
	}
}

func (r *repositoryImpl) Get(ctx context.Context, id string) (*user.User, error) {
	token := ctx.Value(constant.TokenKey).(string)
	u, err := r.keycloakClient.GetUserByID(ctx, token, config.Cfg.KCRealm, id)
	if err != nil {
		return nil, errors.ToDomainErrorFromGoCloak(err)
	}

	return ToDomain(u), nil
}

func (r *repositoryImpl) List(ctx context.Context) ([]*user.User, error) {
	token := ctx.Value(constant.TokenKey).(string)
	enabled := false
	users, err := r.keycloakClient.GetUsers(ctx, token, config.Cfg.KCRealm, gocloak.GetUsersParams{
		Enabled: &enabled,
	})
	if err != nil {
		return nil, errors.ToDomainErrorFromGoCloak(err)
	}

	result := make([]*user.User, len(users))
	for i, u := range users {
		result[i] = ToDomain(u)
	}

	return result, nil
}

func (r *repositoryImpl) Create(ctx context.Context, u *user.User) (*user.User, error) {
	createdUser, err := r.db.CreateUser(ctx, ToCreateUserParams(u))
	if err != nil {
		return nil, errors.ToDomainErrorFromPostgres(err)
	}
	token := ctx.Value(constant.TokenKey).(string)
	user, _ := (r.keycloakClient.GetUserByID(ctx, token, config.Cfg.KCRealm, createdUser.String()))

	return ToDomain(user), nil
}

func (r *repositoryImpl) Update(ctx context.Context, u *user.User) error {
	token := ctx.Value(constant.TokenKey).(string)
	return errors.ToDomainErrorFromPostgres(r.keycloakClient.UpdateUser(ctx, token, config.Cfg.KCRealm, ToUpdateUserParams(u)))
}

func (r *repositoryImpl) Delete(ctx context.Context, id string) error {
	token := ctx.Value(constant.TokenKey).(string)
	return errors.ToDomainErrorFromPostgres(r.keycloakClient.DeleteUser(ctx, token, config.Cfg.KCRealm, id))
}
