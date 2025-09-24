start:
  go run ./cmd/main.go

compose:
  docker compose -f ./docker/db.compose.yaml up

atlas-hash env="local":
  atlas migrate hash --env {{env}}

atlas-gen-migration env="local":
  atlas migrate diff --env {{env}}
