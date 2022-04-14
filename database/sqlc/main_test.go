// go test database/sqlc/main_test.go 

package database

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://godb:123456@localhost:5432/simplebank"
)

var tQuery *Queries

func TestMain(m *testing.M){
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil{
		log.Fatal("Canot connect to database:", err)
	} 
	err1 := testDB.Ping()
	if err1 != nil {
		log.Fatal("Canot connect to database:", err1)
	} else {
		tQuery = New(testDB)
	}
	os.Exit(m.Run())	
}