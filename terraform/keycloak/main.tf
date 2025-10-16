terraform {
  required_version = ">= 1.13.4"
  required_providers {
    keycloak = {
      source  = "keycloak/keycloak"
      version = ">= 5.5.0"
    }
  }
}

resource "keycloak_realm" "electricilies" {
  realm = "electricilies"
}

resource "keycloak_openid_client" "backend" {
  realm_id      = keycloak_realm.electricilies.id
  client_id     = "backend"
  access_type   = "CONFIDENTIAL"
  client_secret = var.backend_client_secret
}

resource "keycloak_openid_client" "frontend" {
  realm_id              = keycloak_realm.electricilies.id
  client_id             = "frontend"
  access_type           = "CONFIDENTIAL"
  client_secret         = var.frontend_client_secret
  standard_flow_enabled = true
  root_url              = var.root_url
  base_url              = var.base_url
  valid_redirect_uris   = ["/*"]
  web_origins           = ["+"]
  admin_url             = var.admin_url
}

locals {
  client_roles = ["admin", "customer", "staff"]
}

resource "keycloak_role" "client_roles" {
  for_each  = toset(local.client_roles)
  realm_id  = keycloak_realm.electricilies.id
  client_id = keycloak_openid_client.frontend.id
  name      = each.key
}
