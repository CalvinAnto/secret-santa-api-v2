package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDatabase() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DBADDR"),
		DBName:               "secret-santa",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Panic(err)
		return
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Panic(pingErr)
		return
	}
	fmt.Println("Connected!")
}

func GetDB() *sql.DB {
	return db
}
