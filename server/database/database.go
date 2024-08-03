package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Create enum with type DBType
type DBType string

const (
	SQLITE     DBType = "SQLITE"
	MYSQL      DBType = "MYSQL"
	POSTGRESQL DBType = "POSTGRESQL"
	MONGODB    DBType = "MONGODB"
)

func GetDatabase() sql.DB {
	log.Println(os.Getenv("DB_TYPE"))

	var db = nil

	switch os.Getenv("DB_TYPE") {
	case string(MONGODB):
	case string(MYSQL):
	case string(POSTGRESQL):
		log.Print("We do not support this yet...\n using sqlite")
	case string(SQLITE):
		db
	default:
		break

	}

	return sqliteDriver
}
