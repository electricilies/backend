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
  <br />
  <a href="https://wakatime.com/badge/github/electricilies/backend">
    <img alt="Wakatime" src="https://wakatime.com/badge/github/electricilies/backend.svg"/>
  </a>
</div>

## Dev

## Resources

- <https://registry.terraform.io/providers/keycloak/keycloak/latest/docs>

### Environment variables

```dotenv
#Backend env var
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_DATABASE=electricilies
DB_PORT=5432 # optional
DB_HOST=localhost
ENV_APP=production # optional, If not set, it will run in development mode
PORT=8080          # optional

# Terraform variable
# Required
TF_VAR_keycloak_terraform_client_secret=BhiJ2qDf9xZp3KrT7LmV5sWe8yA4nC

# Variables with defaults
TF_VAR_keycloak_backend_client_secret=backendclientsecret
TF_VAR_keycloak_frontend_client_secret=frontendclientsecret
TF_VAR_keycloak_root_url=http://localhost:3000
TF_VAR_keycloak_base_url=http://localhost:3000/home
TF_VAR_keycloak_admin_url=/admin
```

### Dev environment

- Tool required are in [mise.toml](./mise.toml) and [flake.nix](./flake.nix)
- Either using `direnv` or [`mise-nix`](https://github.com/mise-plugins/mise-nix) to enable flake
