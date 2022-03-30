package database

import (
	"Service_grpc/conf"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	config := conf.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Dbhost,
		config.DB.Dbport,
		config.DB.Dbname,
		config.DB.Charset)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Could not connect service_grpc")
	}

	return db
}
