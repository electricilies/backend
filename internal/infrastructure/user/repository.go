package user

import (
	"context"
	"encoding/json"

	"backend/config"
	"backend/internal/constant"
	"backend/internal/domain/cart"
	"backend/internal/domain/user"
	"backend/internal/helper"
	"backend/internal/infrastructure/errors"
	"backend/internal/infrastructure/persistence/postgres"
	"backend/pkg/logger"

	"github.com/Nerzal/gocloak/v13"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type repositoryImpl struct {
	db             *postgres.Queries
	s3Client       *s3.Client
	redisClient    *redis.Client
	keycloakClient *gocloak.GoCloak
	tokenManager   helper.TokenManager
	cfg            *config.Config
	logger         *zap.Logger
}

func NewRepository(db *postgres.Queries, s3Client *s3.Client, redisClient *redis.Client, keycloakClient *gocloak.GoCloak, tokenManager helper.TokenManager, cfg *config.Config, logger *zap.Logger) user.Repository {
	return &repositoryImpl{
		db:             db,
		s3Client:       s3Client,
		redisClient:    redisClient,
		keycloakClient: keycloakClient,
		tokenManager:   tokenManager,
		cfg:            cfg,
		logger:         logger,
	}
}

func (r *repositoryImpl) Get(ctx context.Context, id string) (*user.Model, error) {
	cacheKey := constant.UserCachePrefix + id
	cached, err := r.redisClient.Get(ctx, cacheKey).Result()
	switch {
	case err != nil && err != redis.Nil:
		r.logger.Error(constant.ErrRedisGetUserMsg, *logger.CreateRedisGetField(id, cacheKey, err)...)
	default:
		var cachedUser user.Model
		if err := json.Unmarshal([]byte(cached), &cachedUser); err == nil {
			return &cachedUser, nil
		}

	}
	token, err := r.tokenManager.GetClientToken(ctx)
	if err != nil {
		return nil, errors.ToDomainErrorFromGoCloak(err)
	}
	u, err := r.keycloakClient.GetUserByID(ctx, token, r.cfg.KCRealm, id)
	if err != nil {
		return nil, errors.ToDomainErrorFromGoCloak(err)
	}

	domainUser := ToDomain(u)

	if data, err := json.Marshal(domainUser); err == nil {
		r.redisClient.Set(ctx, cacheKey, data, constant.UserCacheTTL)
	}

	return domainUser, nil
}

func (r *repositoryImpl) List(ctx context.Context) ([]*user.Model, error) {
	cached, err := r.redisClient.Get(ctx, constant.UserListCacheKey).Result()
	switch {
	case err != nil && err != redis.Nil:
		r.logger.Error(constant.ErrRedisGetUserMsg, *logger.CreateRedisListField(constant.UserListCacheKey, err)...)
	default:

		var usersCache []*user.Model
		if err := json.Unmarshal([]byte(cached), &usersCache); err == nil {
			return usersCache, nil
		}
	}

	token, err := r.tokenManager.GetClientToken(ctx)
	if err != nil {
		return nil, errors.ToDomainErrorFromGoCloak(err)
	}
	enabled := false
	users, err := r.keycloakClient.GetUsers(ctx, token, r.cfg.KCRealm, gocloak.GetUsersParams{
		Enabled: &enabled,
	})
	if err != nil {
		return nil, errors.ToDomainErrorFromGoCloak(err)
	}

	result := make([]*user.Model, len(users))
	for i, u := range users {
		result[i] = ToDomain(u)
	}

	if data, err := json.Marshal(result); err == nil {
		r.redisClient.Set(ctx, constant.UserListCacheKey, data, constant.UserListCacheTTL)
	}

	return result, nil
}

func (r *repositoryImpl) Create(ctx context.Context, u *user.Model) (*user.Model, error) {
	createdUser, err := r.db.CreateUser(ctx, ToCreateUserParams(u))
	if err != nil {
		return nil, errors.ToDomainErrorFromPostgres(err)
	}
	token, err := r.tokenManager.GetClientToken(ctx)
	if err != nil {
		return nil, errors.ToDomainErrorFromGoCloak(err)
	}
	user, _ := (r.keycloakClient.GetUserByID(ctx, token, r.cfg.KCRealm, createdUser.String()))
	r.redisClient.Del(ctx, constant.UserListCacheKey)

	return ToDomain(user), nil
}

func (r *repositoryImpl) Update(
	ctx context.Context,
	model *user.Model,
	queryParams *user.QueryParams,
) error {
	token, err := r.tokenManager.GetClientToken(ctx)
	if err != nil {
		return errors.ToDomainErrorFromGoCloak(err)
	}
	err = errors.ToDomainErrorFromGoCloak(r.keycloakClient.UpdateUser(ctx, token, r.cfg.KCRealm, *ToUpdateUserParams(model, queryParams.UserID)))
	if err != nil {
		return err
	}

	r.redisClient.Del(ctx, constant.UserCachePrefix+model.ID.String())
	r.redisClient.Del(ctx, constant.UserListCacheKey)

	return nil
}

func (r *repositoryImpl) Delete(ctx context.Context, id string) error {
	token, err := r.tokenManager.GetClientToken(ctx)
	if err != nil {
		return errors.ToDomainErrorFromGoCloak(err)
	}
	err = errors.ToDomainErrorFromGoCloak(r.keycloakClient.DeleteUser(ctx, token, r.cfg.KCRealm, id))
	if err != nil {
		return err
	}

	r.redisClient.Del(ctx, constant.UserCachePrefix+id)
	r.redisClient.Del(ctx, constant.UserListCacheKey)

	return nil
}

// TODO: implement
func (r *repositoryImpl) GetCart(ctx context.Context, id string) (*cart.Model, error) {
	return nil, nil
}
