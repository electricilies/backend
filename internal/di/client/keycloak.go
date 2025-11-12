package client

import (
	"context"
	"log"

	"backend/config"

	"github.com/Nerzal/gocloak/v13"
)

func NewKeycloak(cfg *config.Config) *gocloak.GoCloak {
	client := gocloak.NewClient(cfg.KCBasePath)
	ctx := context.Background()
	token, err := client.LoginClient(ctx, cfg.KCClientId, cfg.KCClientSecret, cfg.KCRealm)
	if err != nil || token == nil {
		log.Printf("error when connecting to keycloak:%s", err)
		return nil
	}

	return client
}
