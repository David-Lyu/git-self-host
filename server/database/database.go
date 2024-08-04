package database

import (
	sqliteDriver "app/database/sqlite"
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

type Database struct {
	Db       *sql.DB
	InsertDB func(string, *sql.DB) (*sql.Stmt, error)
}

var DB *Database

func InitDatabase() {

	if DB == nil {
		log.Println("Runs")

		switch os.Getenv("DB_TYPE") {
		case string(MONGODB):
			fallthrough
		case string(MYSQL):
			fallthrough
		case string(POSTGRESQL):
			log.Print("We do not support this yet...\n Using sqlite")
			fallthrough
		case string(SQLITE):
			fallthrough
		default:
			DB = &Database{
				Db:       sqliteDriver.SqliteDriver(),
				InsertDB: sqliteDriver.InsertTable,
			}
			break
		}

	}
}
