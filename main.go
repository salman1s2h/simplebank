package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import Postgres driver
	"github.com/salman1s2h/simplebank/api"
	db "github.com/salman1s2h/simplebank/db/sqlc"
	"github.com/salman1s2h/simplebank/util"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://admin:admin@localhost:5433/go_db?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

func main() {
	env_obj := util.NewEnv()

	fmt.Printf("Environment: %+v\n", env_obj.AppEnv)

	dbSource := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		env_obj.DBUser, env_obj.DBPass, env_obj.DBHost, env_obj.DBPort, env_obj.DBName,
	)
	conn, err := sql.Open(env_obj.DB_DRIVER, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)

	server, err := api.NewServer(env_obj, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	serverAddr := fmt.Sprintf("%s:%s", env_obj.ServerAddress, env_obj.APP_PORT)

	log.Printf("Starting server at %s", serverAddr)

	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
