provider "keycloak" {
  client_id     = "terraform"
  client_secret = var.terraform_client_secret
  url           = "http://localhost:8080"
}

