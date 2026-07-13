DB_SOURCE=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
postgres: 
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose up
migrateup1:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose up 1
migratedown:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose down
migratedown1:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose down 1
db_docs:
	dbml2sql schema.dbml
db_schema:
	dbml2sql --postgres -o doc/schema.dbml doc/db.dbml
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Nevets9/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1 db_docs db_schema
