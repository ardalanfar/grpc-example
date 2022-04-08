package database

import (
	"database/sql"
	"fmt"
	"grpc-example/conf"
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
		log.Fatal("Could not connect database")
	}
	return db
}
