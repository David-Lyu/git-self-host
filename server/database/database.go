package database

import (
	"log"
	"os"
	
	sqliteDriver "app/database/sqlite"
	"app/models"
	
	_ "github.com/mattn/go-sqlite3"
)

type DBType string

type Databaser interface {
	models.UserDB
}

const (
	SQLITE     DBType = "SQLITE"
	MYSQL      DBType = "MYSQL"
	POSTGRESQL DBType = "POSTGRESQL"
	MONGODB    DBType = "MONGODB"
)

func InitDatabase() Databaser {
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
		return sqliteDriver.NewSqliteDB()
		
	}
}
