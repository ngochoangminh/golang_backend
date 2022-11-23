package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "ngochm"
	password = "ngochm"
	dbname   = "higate"
  )

func ConnectDB() (*sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")
	return db
}