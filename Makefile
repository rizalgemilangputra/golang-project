run-dev:
	go run ./cmd/web/main.go

## make create_migration table_name=users
create_migration:
	migrate create -ext=sql -dir=db/migration -seq ${table_name}

## make migrate_up
migrate_up:
	migrate -path=db/migration -database "postgres://postgres:root@localhost:5432/golang_test?sslmode=disable" -verbose up

## make migrate_up
migrate_down:
	migrate -path=db/migration -database "postgres://postgres:root@localhost:5432/golang_test?sslmode=disable" -verbose down
