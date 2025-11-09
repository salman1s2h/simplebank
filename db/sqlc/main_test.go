package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // Import Postgres driver
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://admin:admin@localhost:5433/go_db?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB
var testStore *Store // ✅ add this line

func TestMain(m *testing.M) {
	var err error
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	fmt.Printf("Connected to db successfully %v", testDB.Stats())
	testQueries = New(testDB)
	testStore = NewStore(testDB) // ✅ initialize store here

	os.Exit(m.Run())
}
