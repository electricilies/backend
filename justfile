gen-migration name:
  migrate -path ./database/schema.sql create -ext sql -dir ./migration/ -seq {{name}}

