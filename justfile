set dotenv-load

main-go := "./cmd/main.go"
bin-out := "./backend"

[doc("Dev build (no optimizations) and run")]
dev:
  go build -gcflags='all=-N -l' -o {{bin-out}} {{main-go}}
  {{bin-out}}

[doc("Dev watch")]
dev-watch:
  air

[doc("Build")]
build:
  go build -o {{bin-out}} {{main-go}}

[doc("Run")]
run: build
  ./{{bin-out}}

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

[doc("Run lint all")]
lint: lint-golangci-lint lint-sqlfluff

[doc("Generate DI wire file")]
gen-wire:
  wire gen internal/di/wire.go

[doc("Generate swagger output to ./docs/")]
gen-swag *args="":
  swag init -g {{main-go}} {{args}}

[doc("Generate sqlc code")]
gen-sqlc *args="":
  sqlc generate {{args}}

[doc("Run gen all")]
gen: gen-wire gen-swag gen-sqlc

format-gofumpt *args="":
  gofumpt -w . {{args}}

format-swag *args="":
  swag fmt {{args}}

format-sqlfluff:
  sqlfluff fix --dialect postgres \
    ./database/ \
    ./database/queries/ \
    ./docker/volume/

[doc("Run format all")]
format: format-gofumpt format-swag format-sqlfluff

check-format-gofumpt *args="":
  gofumpt -l . {{args}}

[doc("Check Format")]
check-format: check-format-gofumpt

[doc("Run gen, static type check, lint, format, suitable for pre-commit")]
pre-commit: gen check-static-type lint format

[doc("Docker compose up")]
compose *args="":
  docker compose -f ./docker/compose.yaml up {{args}}

[private]
[unix]
swagger-local-json:
  echo 'const spec = ' | cat - ./docs/swagger.json > ./docs/swagger-local.js

[unix]
[doc("Generate and open static swagger web locally")]
swagger-web: gen-swag swagger-local-json
  xdg-open ./docs/local.html

db-seed-fake:
  psql postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE} -f ./database/seed-fake.sql

[doc("Apply schema for local development")]
atlas-apply-schema env="local" *args='':
  atlas schema apply --env {{env}} {{args}}

atlas-gen-migration env="dev":
  atlas migrate diff --env {{env}}

[doc("Export a Keycloak realm to JSON")]
export-realm container="electricilies-backend-keycloak-1" realm="electricilies":
  docker exec \
    -it \
    {{container}} \
    /opt/keycloak/bin/kc.sh export \
    --optimized \
    --realm {{realm}} \
    --file /opt/keycloak/{{realm}}.json || true
  docker cp {{container}}:/opt/keycloak/{{realm}}.json ./keycloak/{{realm}}-export.json

[doc("Import a Keycloak realm from JSON")]
import-realm container="electricilies-backend-keycloak-1" file="./keycloak/electricilies-export.json" realm="electricilies":
  docker cp {{file}} {{container}}:/opt/keycloak/{{realm}}-export.json
  docker exec \
    -it \
    {{container}} \
    /opt/keycloak/bin/kc.sh import \
    --optimized \
    --file /opt/keycloak/{{realm}}-export.json

gen-ctags:
  ctags -R \
    --languages=Go \
    --exclude=.git \
    --exclude=terraform \
    --exclude=http \
    --exclude=migration \
    --exclude=database \
    --exclude=build \
    --exclude=docker \
    --exclude=*.toml \
    --exclude=vendor
