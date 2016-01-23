package database

import (
	"fmt"
	"github.com/dabfleming/foo/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

func init() {
	var err error

	// TODO unchecked type assertions, BAD!
	connectionString := fmt.Sprintf(
		"user=%s host=%s port=%s password=%s dbname=%s sslmode=%s",
		config.Get().Database["user"].(string),
		config.Get().Database["host"].(string),
		config.Get().Database["port"].(string),
		config.Get().Database["password"].(string),
		config.Get().Database["dbname"].(string),
		config.Get().Database["sslmode"].(string),
	)

	db, err = sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Could not connect to database. Error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping database. Error: %v", err)
	}
}