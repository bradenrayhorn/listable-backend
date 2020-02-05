### Listable Backend

The backend for a shopping list app. Written in Go.

### Setup

Migrations are handled by [migrate](https://github.com/golang-migrate/migrate). Install migrate, then run the migrations.

Example command: `migrate -database "mysql://root:password@tcp(127.0.0.1:3306)/listable"  -path ./db/migrations/ up`
