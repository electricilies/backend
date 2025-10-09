variable "temp_db" {
  type    = string
  default = "docker://postgres/17.6-alpine3.22/dev"
}

variable "db_url" {
  type    = string
  default = "postgres://${getenv("DB_USERNAME")}:${getenv("DB_PASSWORD")}@${getenv("DB_HOST")}:${getenv("DB_PORT")}/${getenv("DB_DATABASE")}?sslmode=disable"
}

locals {
  schema_path    = "file://database/schema.sql"
  migration_path = "file://migration"
}

env "local" {
  src     = local.schema_path
  url     = var.db_url
  dev     = var.temp_db
  schemas = ["public"]
  migration {
    dir = local.migration_path
  }
}

env "dev" {
  src     = local.schema_path
  url     = var.db_url
  schemas = ["public"]
  migration {
    dir = local.migration_path
  }
}
