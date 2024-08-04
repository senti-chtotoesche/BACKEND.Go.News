package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func DbInit() {
	var err error
	DB, err = sql.Open("postgres",
		"host=localhost"+
			" port=5432 "+
			"user=******"+
			" password=*********"+
			" dbname=******* "+
			"sslmode=disable")

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
}
