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
  realm                          = "electricilies"
  access_code_lifespan           = "1h"
  registration_allowed           = true
  registration_email_as_username = true
  reset_password_allowed         = true
  remember_me                    = true
  login_with_email_allowed       = true
  attributes = {
    userProfileEnable = true
  }
}

resource "keycloak_openid_client" "backend" {
  realm_id      = keycloak_realm.electricilies.id
  client_id     = "backend"
  access_type   = "CONFIDENTIAL"
  client_secret = var.keycloak_backend_client_secret
}

resource "keycloak_openid_client" "frontend" {
  realm_id                        = keycloak_realm.electricilies.id
  client_id                       = "frontend"
  access_type                     = "CONFIDENTIAL"
  client_secret                   = var.keycloak_frontend_client_secret
  standard_flow_enabled           = true
  standard_token_exchange_enabled = true
  direct_access_grants_enabled    = true
  root_url                        = var.keycloak_frontend_root_url
  base_url                        = var.keycloak_frontend_base_url
  valid_redirect_uris             = ["*"]
  web_origins                     = ["+"]
  admin_url                       = var.keycloak_frontend_admin_url
}

resource "keycloak_realm_user_profile" "userprofile" {
  realm_id = keycloak.electricilies.id
  attribute {
    name         = "phone_number"
    display_name = "Phone Number"
  }
}

locals {
  map = {
    admin    = "admin"
    customer = "customer"
    staff    = "staff"
  }
}

resource "keycloak_role" "client_roles" {
  for_each  = toset(local.map)
  realm_id  = keycloak_realm.electricilies.id
  client_id = keycloak_openid_client.frontend.id
  name      = each.value
}

resource "keycloak_default_roles" "default_roles" {
  realm_id      = keycloak_realm.electricilies.id
  default_roles = [local.map.customer]
}


resource "keycloak_user" "user" {
  for_each   = local.map
  realm_id   = keycloak_realm.realm.id
  username   = each.key
  enabled    = true
  email      = "${each.key}@example.com"
  first_name = title(each.value)
  last_name  = "User"
}

