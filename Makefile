DB_URL=postgresql://admin:admin@localhost:5433/go_db?sslmode=disable

createdb:
	docker exec -it postgres_container createdb -U postgres go_db

dropdb:
	docker exec -it postgres_container dropdb -U postgres go_db

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up


migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v ./...


server:
	go run main.go

.PHONY: network postgres createdb dropdb migrateup migratedown sqlc test server