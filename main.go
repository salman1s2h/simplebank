package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Import Postgres driver
	"github.com/salman1s2h/simplebank/api"
	db "github.com/salman1s2h/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://admin:admin@localhost:5433/go_db?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)

	server := api.NewServer(store)
	log.Printf("Starting server at %s", serverAddress)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
