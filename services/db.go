package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // sql driver
)

// GetDB : db init
func GetDB() (*sql.DB, error) {
	dbStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"), os.Getenv("password"),
		os.Getenv("host"), os.Getenv("port"), os.Getenv("database"),
	)
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		fmt.Println(err)
	}

	return db, err
}
