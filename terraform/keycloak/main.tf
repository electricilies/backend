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
  realm_id    = keycloak_realm.electricilies.id
  client_id   = "backend"
  access_type = "CONFIDENTIAL"
}

resource "keycloak_openid_client" "frontend" {
  realm_id    = keycloak_realm.electricilies.id
  client_id   = "frontend"
  access_type = "CONFIDENTIAL"
}
