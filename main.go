package main

import (
	"database/sql"
	"log"

	"github.com/codewithtoucans/simplebank/api"
	db "github.com/codewithtoucans/simplebank/db/sqlc"
	"github.com/codewithtoucans/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(config, store)

	log.Printf("Starting server at %s", config.ServerAddress)
	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server", err)
		return
	}
}
