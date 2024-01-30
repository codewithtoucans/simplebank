postgres:
	docker run --name mypostgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=rootgzy -d postgres

stoppg:
	docker stop mypostgres

startpg:
	docker start mypostgres

restartpg:
	docker restart mypostgres

psql:
	docker exec -it mypostgres psql -U root -d simple_bank

createdb:
	docker exec -it mypostgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it mypostgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:rootgzy@localhost:5432/simple_bank?sslmode=disable" -verbose up
	
migrateup1:
	migrate -path db/migration -database "postgresql://root:rootgzy@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:rootgzy@localhost:5432/simple_bank?sslmode=disable" -verbose down
	
migratedown1:
	migrate -path db/migration -database "postgresql://root:rootgzy@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -cover -v -count=1 ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown
