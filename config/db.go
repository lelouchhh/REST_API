package config

import (
	"controllers"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

// Connecting to db
func Connect() *sql.DB {
	/*opts := &pg.Options{
		User: "inotech",
		Password: "platex",
		Addr: "178.154.254.105:5432",
		Database: "platex",
	}*/
	connStr := "host=178.154.254.105 user=inotech password=platex dbname=platex sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Printf("Connected to db")
	controllers.InitiateDB(db)
	return db
}