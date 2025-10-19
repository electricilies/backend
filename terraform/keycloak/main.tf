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
  registration_allowed     = true
  reset_password_allowed   = true
  remember_me              = true
  login_with_email_allowed = true
  attributes = {
    userProfileEnable = true
  }
}

resource "keycloak_openid_client" "backend" {
  realm_id      = keycloak_realm.electricilies.id
  client_id     = "backend"
  name          = "Backend"
  access_type   = "CONFIDENTIAL"
  client_secret = var.keycloak_backend_client_secret
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
    name         = "birthday"
    display_name = "Birthday"
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
    name         = "deleted-at"
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

locals {
  roles = [
    "admin",
    "staff",
    "customer"
  ]

  users = {
    admin = {
      password = "admin",
      role     = "admin"
    },
    staff = {
      password = "staff",
      role     = "staff",
    },
    customer = {
      password = "customer",
      role     = "customer"
    },
  }
}

resource "keycloak_role" "backend_roles" {
  for_each = toset(local.roles)

  realm_id  = keycloak_realm.electricilies.id
  client_id = keycloak_openid_client.backend.id
  name      = each.key
}

resource "keycloak_default_roles" "default_roles" {
  realm_id = keycloak_realm.electricilies.id
  default_roles = [
    "${keycloak_openid_client.backend.client_id}/${keycloak_role.backend_roles["customer"].name}",
  ]
}

resource "keycloak_user" "users" {
  for_each = local.users
  depends_on = [
    keycloak_realm_user_profile.userprofile,
  ]

  realm_id = keycloak_realm.electricilies.id
  username = each.key
  initial_password {
    value = each.value.password
  }
  email          = "${each.key}@example.com"
  email_verified = true
  attributes = {
    first_name = title(each.key)
  }
}

resource "keycloak_user_roles" "users" {
  for_each = local.users

  realm_id = keycloak_realm.electricilies.id
  user_id  = keycloak_user.users[each.key].id
  role_ids = [
    keycloak_role.backend_roles[each.value.role].id,
  ]
}

