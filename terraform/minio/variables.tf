variable "minio_server_endpoint" {
  type    = string
  default = "localhost:9000"
}

variable "minio_username" {
  type    = string
  default = "electricilies"
}

variable "minio_password" {
  type    = string
  default = "electricilies"
}

variable "backend_webhook" {
  type    = string
  default = "http://localhost:8080/api"
}
