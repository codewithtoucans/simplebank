package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:rootgzy@localhost:5432/simple_bank?sslmode=disable"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer testDB.Close()

	if err = testDB.Ping(); err != nil {
		log.Fatal("cannot ping to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
