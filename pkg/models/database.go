package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB
var dbErr error

// SetupDatabase setup and test database connection
func SetupDatabase() {
	db, dbErr = sql.Open("mysql", toDatabaseDSN())
	if dbErr != nil {
		log.Fatalln("unable to connect to database: ", dbErr.Error())
		return
	}

	dbErr = db.Ping()
	if dbErr != nil {
		log.Fatalln("unable to ping database: ", dbErr.Error())
		return
	}

	createDatabaseTable()
}

// createDatabaseTable Create the database table that will hold our data
func createDatabaseTable() {
	query := `CREATE TABLE IF NOT EXISTS todos (
		id VARCHAR(255) NOT NULL UNIQUE,
		task VARCHAR(255) NOT NULL,
		completed TINYINT(1) NOT NULL DEFAULT 0,

		primary key(id)
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalln("unable to create database table: ", err.Error())
	}
}

func toDatabaseDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
}
