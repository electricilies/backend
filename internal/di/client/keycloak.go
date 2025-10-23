package client

import (
	"backend/config"
	"context"
	"log"

	"github.com/Nerzal/gocloak/v13"
)

func NewKeycloak() *gocloak.GoCloak {
	client := gocloak.NewClient(config.Cfg.KCBasePath)
	ctx := context.Background()
	token, err := client.LoginClient(ctx, config.Cfg.KCClientID, config.Cfg.KCClientSecret, config.Cfg.KCRealm)
	if err != nil || token == nil {
		log.Printf("error when connecting to keycloak:%s", err)
		return nil
	}

	return client
}
