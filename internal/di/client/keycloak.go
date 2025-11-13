package client

import (
	"context"
	"log"

	"backend/config"

	"github.com/Nerzal/gocloak/v13"
)

func NewKeycloak(ctx context.Context, cfg *config.Config) *gocloak.GoCloak {
	client := gocloak.NewClient(cfg.KCBasePath)
	token, err := client.LoginClient(ctx, cfg.KCClientId, cfg.KCClientSecret, cfg.KCRealm)
	if err != nil || token == nil {
		log.Printf("error when connecting to keycloak:%s", err)
		return nil
	}

	return client
}
