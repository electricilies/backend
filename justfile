start:
  go run ./cmd/main.go

debug:
  dlv debug --headless --listen=:4444 ./cmd/main.go

compose:
  docker compose -f ./docker/db.compose.yaml up

atlas-apply-schema env="local":
  atlas schema apply --env {{env}}

atlas-gen-migration env="dev":
  atlas migrate diff --env {{env}}
