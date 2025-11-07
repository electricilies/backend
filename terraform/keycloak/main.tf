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
  realm                    = "electricilies"
  access_code_lifespan     = "12h"
  access_token_lifespan    = "8760h"
  duplicate_emails_allowed = false
  login_with_email_allowed = true
  registration_allowed     = true
  remember_me              = true
  reset_password_allowed   = true
  verify_email             = false
  attributes = {
    userProfileEnable = true
  }
}

resource "keycloak_openid_client" "backend" {
  realm_id                 = keycloak_realm.electricilies.id
  client_id                = "backend"
  name                     = "Backend"
  access_type              = "CONFIDENTIAL"
  client_secret            = var.keycloak_backend_client_secret
  service_accounts_enabled = true
}

resource "keycloak_openid_client" "frontend" {
  realm_id                        = keycloak_realm.electricilies.id
  client_id                       = "frontend"
  access_type                     = "CONFIDENTIAL"
  name                            = "Frontend"
  client_secret                   = var.keycloak_frontend_client_secret
  standard_flow_enabled           = true
  standard_token_exchange_enabled = true
  direct_access_grants_enabled    = true
  root_url                        = var.keycloak_frontend_root_url
  base_url                        = var.keycloak_frontend_base_url
  valid_redirect_uris             = ["*"]
  web_origins                     = ["*"]
  admin_url                       = var.keycloak_frontend_admin_url
}

resource "keycloak_realm_user_profile" "userprofile" {
  realm_id = keycloak_realm.electricilies.id

  attribute {
    name = "username"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "first_name"
    display_name = "First Name"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "last_name"
    display_name = "Last Name"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name = "email"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "phone_number"
    display_name = "Phone Number"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    validator {
      name = "pattern"
      config = {
        pattern = "^0[0-9]{9,10}$"
      }
    }
    annotations = {
      inputType = "html5-tel"
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "address"
    display_name = "Address"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "date_of_birth"
    display_name = "Date of Birth"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    annotations = {
      inputType = "html5-date"
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "role"
    display_name = "Role"
    permissions {
      view = ["admin", "user"]
      edit = ["admin"]
    }
    annotations = {
      inputType = "select"
    }
    validator {
      name = "options"
      config = {
        options = jsonencode(
          [
            "admin",
            "staff",
            "customer",
          ],
        )
      }
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "deleted_at"
    display_name = "Deleted At"
    permissions {
      view = ["admin"]
      edit = ["admin"]
    }
    annotations = {
      inputType = "html5-datetime-local"
    }
  }
}

resource "keycloak_openid_client_scope" "role" {
  realm_id               = keycloak_realm.electricilies.id
  name                   = "role"
  description            = "Role of Electricilies app"
  include_in_token_scope = true
}

resource "keycloak_generic_protocol_mapper" "role" {
  name            = "role"
  realm_id        = keycloak_realm.electricilies.id
  client_scope_id = keycloak_openid_client_scope.role.id
  protocol        = "openid-connect"
  protocol_mapper = "oidc-usermodel-attribute-mapper"
  config = {
    "introspection.token.claim" : "true",
    "userinfo.token.claim" : "true",
    "user.attribute" : "role",
    "id.token.claim" : "false",
    "lightweight.claim" : "false",
    "access.token.claim" : "true",
    "claim.name" : "role",
    "jsonType.label" : "String"
  }
}

resource "keycloak_openid_client_default_scopes" "frontend" {
  realm_id  = keycloak_realm.electricilies.id
  client_id = keycloak_openid_client.frontend.id

  default_scopes = [
    "profile",
    "email",
    "basic",
    keycloak_openid_client_scope.role.name,
    # Below seem we don't need it
    # "acr",
    # "roles",
    # "service_account",
    # "web-origins",
  ]
}

data "keycloak_openid_client" "account" {
  realm_id  = keycloak_realm.electricilies.id
  client_id = "account"
}

resource "keycloak_openid_client_service_account_role" "backend" {
  realm_id                = keycloak_realm.electricilies.id
  client_id               = data.keycloak_openid_client.account.id
  service_account_user_id = keycloak_openid_client.backend.service_account_user_id
  role                    = "view-profile"
}

locals {
  users = {
    admin = {
      password      = "admin",
      role          = "admin"
      first_name    = "admin",
      last_name     = "admin"
      email         = "admin@example.com"
      phone_number  = "0909909909"
      address       = "admin address"
      date_of_birth = "01/01/2001"
    },
    staff = {
      password      = "staff",
      role          = "staff",
      first_name    = "staff",
      last_name     = "staff"
      email         = "staff@example.com"
      phone_number  = "0909909909"
      address       = "staff address"
      date_of_birth = "01/01/2001"
    },
    customer = {
      password      = "customer",
      role          = "customer"
      first_name    = "customer",
      last_name     = "customer"
      email         = "customer@example.com"
      phone_number  = "0909909909"
      address       = "customer address"
      date_of_birth = "01/01/2001"
    },
  }
}

resource "keycloak_user" "users" {
  for_each = local.users
  depends_on = [
    keycloak_realm_user_profile.userprofile,
  ]

  realm_id = keycloak_realm.electricilies.id
  username = each.key
  initial_password {
    value     = each.value.password
    temporary = false
  }
  email          = "${each.key}@example.com"
  email_verified = true
  attributes = {
    first_name    = each.value.first_name
    role          = each.value.role,
    first_name    = each.value.first_name,
    last_name     = each.value.last_name,
    phone_number  = each.value.phone_number,
    address       = each.value.address,
    date_of_birth = each.value.date_of_birth,
  }
}
