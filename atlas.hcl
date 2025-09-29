variable "temp_db" {
  type    = string
  default = "docker://postgres/18.0-trixie/dev"
}

variable "local_db_url" {
  type    = string
  default = "postgres://${getenv("DB_USERNAME")}:${getenv("DB_PASSWORD")}@${getenv("DB_HOST")}:${getenv("DB_PORT")}/${getenv("DB_DATABASE")}?sslmode=disable"
}

locals {
  schema_path    = "file://database/schema.sql"
  migration_path = "file://migration"
}

env "local" {
  src = local.schema_path
  url = var.local_db_url
  dev = var.temp_db
}

env "dev" {
  src = local.schema_path
  dev = var.temp_db
  migration {
    dir = local.migration_path
  }
}
