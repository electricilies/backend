start:
  go run ./cmd/main.go

build:
  go build -o backend ./cmd/main.go

debug:
  dlv debug --headless --listen=:4444 ./cmd/main.go

compose:
  docker compose -f ./docker/compose.yaml up

swagger-docs:
  swag init -g ./cmd/main.go

atlas-apply-schema env="local" *args='':
  atlas schema apply --env {{env}} {{args}}

atlas-gen-migration env="dev":
  atlas migrate diff --env {{env}}
