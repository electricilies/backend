package helper

import (
	"context"

	"backend/config"

	"github.com/Nerzal/gocloak/v13"
)

type TokenManager interface {
	GetClientToken(context.Context) (string, error)
}

type tokenManager struct {
	keycloakClient *gocloak.GoCloak
	token          *gocloak.JWT
}

func NewTokenManager(keycloakClient *gocloak.GoCloak) TokenManager {
	return &tokenManager{
		keycloakClient: keycloakClient,
	}
}

func (tm *tokenManager) GetClientToken(ctx context.Context) (string, error) {
	if tm.token == nil || tm.token.ExpiresIn < 10 {
		return tm.token.AccessToken, nil
	}
	token, err := tm.keycloakClient.LoginClient(ctx, config.Cfg.KcClientId, config.Cfg.KcClientSecret, config.Cfg.KcRealm)
	if err != nil {
		return "", err
	}
	tm.token = token
	return token.AccessToken, nil
}
