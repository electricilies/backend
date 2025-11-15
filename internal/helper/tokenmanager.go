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
	srvCfg         *config.Server
}

func NewTokenManager(keycloakClient *gocloak.GoCloak, srvCfg *config.Server) TokenManager {
	return &tokenManager{
		keycloakClient: keycloakClient,
		srvCfg:         srvCfg,
	}
}

func (tm *tokenManager) GetClientToken(ctx context.Context) (string, error) {
	if tm.token == nil || tm.token.ExpiresIn < 10 {
		return tm.token.AccessToken, nil
	}
	token, err := tm.keycloakClient.LoginClient(ctx, tm.srvCfg.KCClientId, tm.srvCfg.KCClientSecret, tm.srvCfg.KCRealm)
	if err != nil {
		return "", err
	}
	tm.token = token
	return token.AccessToken, nil
}
