set dotenv-load

main-go := "./cmd/main.go"
bin-out := "./backend"

[doc("Watch dev")]
dev:
  air

[doc("Build")]
build:
  go build -o {{bin-out}} {{main-go}}

[doc("Debug")]
debug:
  dlv debug --headless --listen=:4444 {{main-go}}

[doc("Run test")]
test *args="":
  go test ./... {{args}}

check-static-type:
  go vet ./cmd/main.go

lint-golangci-lint *args="":
  golangci-lint run {{args}}

lint-sqlfluff:
  sqlfluff lint --dialect postgres \
    ./database/ \
    ./database/queries/ \
    ./docker/volume/

[doc("Run lint")]
lint: lint-golangci-lint lint-sqlfluff

format-gofumpt *args="":
  gofumpt -w . {{args}}

format-sqlfluff:
  sqlfluff fix --dialect postgres \
    ./database/ \
    ./database/queries/ \
    ./docker/volume/

[doc("Run Format")]
format: format-gofumpt format-sqlfluff

check-format-gofumpt *args="":
  gofumpt -l . {{args}}

[doc("Check Format")]
check-format: check-format-gofumpt

[doc("Docker compose up")]
compose *args="":
  docker compose -f ./docker/compose.yaml up {{args}}

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

[doc("Export a Keycloak realm to JSON")]
export-realm container="electricilies-backend-keycloak-1" realm="electricilies":
  docker exec -it {{container}} \
    /opt/keycloak/bin/kc.sh export \
    --realm {{realm}} \
    --file /{{realm}}.json
  docker cp {{container}}:/{{realm}}.json ./keycloak/{{realm}}-export.json

[doc("Import a Keycloak realm from JSON")]
import-realm container="electricilies-backend-keycloak-1" file="./keycloak/electricilies-export.json" realm="electricilies":
  docker cp {{file}} {{container}}:/{{realm}}-export.json
  docker exec -it {{container}} \
    /opt/keycloak/bin/kc.sh import \
    --file {{realm}}-export.json
