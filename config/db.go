package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "thien123",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected Database!")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	fmt.Println("Disconnect Database!")
	db.Close()
}
