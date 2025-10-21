package client

import (
	"backend/config"

	"github.com/Nerzal/gocloak/v13"
)

func NewKeycloak() *gocloak.GoCloak {
	return gocloak.NewClient(config.Cfg.KCBasePath)
}
