# BACKEND

<div align=center>
  <a href="https://sonarcloud.io/summary/new_code?id=electricilies.backend">
    <img alt="SonarQube Quality Gate" src="https://sonarcloud.io/api/project_badges/measure?project=electricilies.backend&metric=alert_status"/>
  </a>
  <a href="https://sonarcloud.io/summary/new_code?id=electricilies.backend">
    <img alt="SonarQube Quality Bug" src="https://sonarcloud.io/api/project_badges/measure?project=electricilies.backend&metric=bugs"/>
  </a>
  <a href="https://sonarcloud.io/summary/new_code?id=electricilies.backend">
    <img alt="SonarQube Quality Code Smells" src="https://sonarcloud.io/api/project_badges/measure?project=electricilies.backend&metric=code_smells"/>
  </a>
  <a href="https://sonarcloud.io/summary/new_code?id=electricilies.backend">
    <img alt="SonarQube Quality Maintainability Rating" src="https://sonarcloud.io/api/project_badges/measure?project=electricilies.backend&metric=sqale_rating"/>
  </a>
  <br />
  <a href="https://wakatime.com/badge/github/electricilies/backend">
    <img alt="Wakatime" src="https://wakatime.com/badge/github/electricilies/backend.svg"/>
  </a>
</div>

## Enviroment variables

```env
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_DATABASE=electricilies
DB_PORT=5432
DB_HOST=localhost
```

## Build and Run

```bash
go build -o backend cmd/main.go
./backend
```
