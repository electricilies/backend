variable "terraform_client_secret" {
  type = string
}

variable "backend_client_secret" {
  type    = string
  default = "backendclientsecret"
}

variable "frontend_client_secret" {
  type    = string
  default = "frontendclientsecret"
}

variable "root_url" {
  type    = string
  default = "http://localhost:3000"
}

variable "base_url" {
  type    = string
  default = "http://localhost:3000/home"
}

variable "admin_url" {
  type    = string
  default = "/admin"
}

