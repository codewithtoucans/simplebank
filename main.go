package main

import (
	"database/sql"
	"log"

	"github.com/codewithtoucans/simplebank/api"
	"github.com/codewithtoucans/simplebank/config"
	db "github.com/codewithtoucans/simplebank/db/sqlc"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
