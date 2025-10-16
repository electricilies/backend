set dotenv-load

main-go := "./cmd/main.go"
bin-out := "./backend"

[doc("Watch dev")]
dev:
  mise watch dev

[doc("Build")]
build:
  go build -o {{bin-out}} {{main-go}}

[doc("Debug")]
debug:
  dlv debug --headless --listen=:4444 {{main-go}}

[doc("Docker compose up")]
compose:
  docker compose -f ./docker/compose.yaml up

[doc("Generate swagger output to ./docs/ with swag")]
swagger-docs *args="":
  swag init -g {{main-go}} {{args}}

[private]
[unix]
swagger-local-json:
  echo 'const spec = ' | cat - ./docs/swagger.json > ./docs/swagger-local.js

[unix]
[doc("Generate and open static swagger web locally")]
swagger-web: swagger-docs swagger-local-json
  xdg-open ./docs/local.html

[doc("Apply schema for local development")]
atlas-apply-schema env="local" *args='':
  atlas schema apply --env {{env}} {{args}}

atlas-gen-migration env="dev":
  atlas migrate diff --env {{env}}
