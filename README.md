# BACKEND

<div align=center>
  <a href="https://sonarcloud.io/summary/new_code?id=electricilies_backend">
    <img alt="SonarQube Quality Gate" src="https://sonarcloud.io/api/project_badges/measure?project=electricilies_backend&metric=alert_status"/>
  </a>
  <a href="https://sonarcloud.io/summary/new_code?id=electricilies_backend">
    <img alt="SonarQube Quality Bug" src="https://sonarcloud.io/api/project_badges/measure?project=electricilies_backend&metric=bugs"/>
  </a>
  <a href="https://sonarcloud.io/summary/new_code?id=electricilies_backend">
    <img alt="SonarQube Quality Code Smells" src="https://sonarcloud.io/api/project_badges/measure?project=electricilies_backend&metric=code_smells"/>
  </a>
  <a href="https://sonarcloud.io/summary/new_code?id=electricilies_backend">
    <img alt="SonarQube Quality Maintainability Rating" src="https://sonarcloud.io/api/project_badges/measure?project=electricilies_backend&metric=sqale_rating"/>
  </a>
  <a href="https://codecov.io/gh/electricilies/backend">
    <img alt="Codecov" src="https://codecov.io/gh/electricilies/backend/branch/main/graph/badge.svg"/>
  </a>
  <br />
  <a href="https://wakatime.com/badge/github/electricilies/backend">
    <img alt="Wakatime" src="https://wakatime.com/badge/github/electricilies/backend.svg"/>
  </a>
</div>

## Dev

### Environment variables

```dotenv
# Backend env var
DB_USERNAME=electricilies
DB_PASSWORD=electricilies
DB_DATABASE=electricilies
DB_PORT=5432 # optional
DB_HOST=localhost
PORT=8080                 # optional
TIMEZONE=Asia/Ho_Chi_Minh # optional

# Keycloak
KC_CLIENT_ID=backend
KC_CLIENT_SECRET=electricilies
KC_REALM=electricilies
KC_BASE_PATH=http://localhost:8081

# S3 / S3 Compatible (MinIO)
S3_ACCESS_KEY=electricilies
S3_SECRET_KEY=electricilies
S3_REGION_NAME=us-east-1
S3_ENDPOINT=http://localhost:9000
S3_BUCKET=electricilies

# Redis / Redis Compatible
REDIS_ADDRESS=localhost:6379

# Terraform variable
TF_VAR_keycloak_terraform_client_secret= # Create manually in the UI from keycloak terraform docs
```

> [!NOTE]
>
> - More extra/optional:
>   - [./docker/compose.yaml](./docker/compose.yaml)
>   - [./terraform/keycloak/variables.tf](./terraform/keycloak/variables.tf)

### Dev environment

- Tool required are in [mise.toml](./mise.toml)

### Note

- Running gen first, then format

## References

- keycloak
  - <https://registry.terraform.io/providers/keycloak/keycloak/latest/docs>
- sqlc
  - <https://github.com/sqlc-dev/sqlc/discussions/364>
  - <https://github.com/sqlc-dev/sqlc/issues/2061>
    <https://github.com/coder/coder/tree/main/coderd/database/queries>
    <https://github.com/riverqueue/river/tree/master/riverdriver/riverpgxv5/internal/dbsqlc>
