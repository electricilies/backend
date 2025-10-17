provider "keycloak" {
  client_id     = "terraform"
  client_secret = var.keycloak_terraform_client_secret
  url           = "http://localhost:8081"
}

