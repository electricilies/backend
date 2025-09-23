env "local" {
  src = "file://database/schema.sql"
  dev = "docker://postgres/17.6-trixie/dev?search_path=public"
  migration {
    dir = "file://migration"
  }
}
