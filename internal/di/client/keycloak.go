package client

import (
	"context"
	"log"

	"backend/config"

	"github.com/Nerzal/gocloak/v13"
)

func NewKeycloak(cfg *config.Config) *gocloak.GoCloak {
	client := gocloak.NewClient(cfg.KcBasePath)
	ctx := context.Background()
	token, err := client.LoginClient(ctx, cfg.KcClientId, cfg.KcClientSecret, cfg.KcRealm)
	if err != nil || token == nil {
		log.Printf("error when connecting to keycloak:%s", err)
		return nil
	}

	return client
}
