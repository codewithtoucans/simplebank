package main

import (
	"database/sql"
	"log"

	"github.com/codewithtoucans/simplebank/api"
	db "github.com/codewithtoucans/simplebank/db/sqlc"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dbDriver      = "pgx"
	dbSource      = "postgresql://root:rootgzy@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(serverAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
