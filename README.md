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
    <img alt="Codecov Business Logic" src="https://img.shields.io/codecov/c/github/electricilies/backend?component=business-logic&label=coverage%20business%20logic"/>
  </a>
  <a href="https://codecov.io/gh/electricilies/backend">
    <img alt="Codecov Application" src="https://img.shields.io/codecov/c/github/electricilies/backend?component=application&label=coverage%20application"/>
  </a>
  <br />
  <a href="https://wakatime.com/badge/github/electricilies/backend">
    <img alt="Wakatime" src="https://wakatime.com/badge/github/electricilies/backend.svg"/>
  </a>
</div>

## Dev

### Environment variables

Read at [mise.toml](./mise.toml)

> [!NOTE]
>
> - More extra/optional:
>   - [./docker/compose.yaml](./docker/compose.yaml)
>   - [./terraform/keycloak/variables.tf](./terraform/keycloak/variables.tf)
>   - [./terraform/minio/variables.tf](./terraform/minio/variables.tf)

### Dev environment

- Tool required are in [mise.toml](./mise.toml)

### Bootstrap

1. `just compose` to spin up all services
   > with `db` service, it init, create schema, trigger, and other necessary stuff
2. `just` to run backend
3. Setup keycloak
   1. Import `terraform` client from ./keycloak/master-terraform-client-export.json
   2. Add `Service account roles` `realm/admin` for `terraform` client
   3. `cd ./terraform/keycloak/`
   4. `terraform apply -auto-aprove`
4. MinIO
   1. `cd ./terraform/minio/`
   2. `terraform apply -auto-aprove`
5. `just db-seed` to seed data to backend database (optional)

### Note

- Running gen first, then format
- Keycloak endpoint:
  - [well-known](http://localhost:8081/realms/electricilies/.well-known/openid-configuration)

## References

- keycloak
  - Auth with keycloak.app: <https://www.keycloak.org/app/?url=http://localhost:8081&realm=electricilies&client=frontend>
  - <https://registry.terraform.io/providers/keycloak/keycloak/latest/docs>
- MinIO
  - Bucket notification: <https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/>
- sqlc
  - <https://github.com/sqlc-dev/sqlc/discussions/364>
  - <https://github.com/sqlc-dev/sqlc/issues/2061>
  - <https://github.com/coder/coder/tree/main/coderd/database/queries>
  - <https://github.com/riverqueue/river/tree/master/riverdriver/riverpgxv5/internal/dbsqlc>
  - <https://gist.github.com/juliogreff/88e585fed5d710044d69f4eca7bf1cb7>
