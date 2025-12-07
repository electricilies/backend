variable "minio_server_endpoint" {
  type = string
}

variable "minio_ssl" {
  type = bool
}

variable "minio_username" {
  type = string
}

variable "minio_password" {
  type      = string
  sensitive = true
}
