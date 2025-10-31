package client

import (
	"context"
	"log"

	"backend/config"

	"github.com/Nerzal/gocloak/v13"
)

func NewKeycloak() *gocloak.GoCloak {
	client := gocloak.NewClient(config.Cfg.KcBasePath)
	ctx := context.Background()
	token, err := client.LoginClient(ctx, config.Cfg.KcClientId, config.Cfg.KcClientSecret, config.Cfg.KcRealm)
	if err != nil || token == nil {
		log.Printf("error when connecting to keycloak:%s", err)
		return nil
	}

	return client
}
