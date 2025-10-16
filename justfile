main-go := "./cmd/main.go"
bin-out := "./backend"

dev:
  mise watch dev

build:
  go build -o {{bin-out}} {{main-go}}

debug:
  dlv debug --headless --listen=:4444 {{main-go}}

compose:
  docker compose -f ./docker/compose.yaml up

swagger-docs:
  swag init -g {{main-go}}

[unix]
swagger-web: swagger-docs
  #!/usr/bin/env bash
  set -euo pipefail
  sleep 1 && xdg-open http://localhost:8081/local.html &
  python3 -m http.server 8081 --directory=./docs/

atlas-apply-schema env="local" *args='':
  atlas schema apply --env {{env}} {{args}}

atlas-gen-migration env="dev":
  atlas migrate diff --env {{env}}
