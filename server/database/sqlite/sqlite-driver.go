package sqliteDriver

import (
	"database/sql"
	"log"
	"os"
	
	_ "github.com/mattn/go-sqlite3"
)

/*
 
 */

type SqliteDB struct {
	db *sql.DB
}

func (sd SqliteDB) InitDb() {
	sd.createUserTable()
	sd.createRepoTable()
}

func NewSqliteDB() SqliteDB {
	if err := os.Remove("./database/sqlite/sqlite-database.db"); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("./database/sqlite/sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")
	
	sqliteDatabase, err := sql.Open("sqlite3", "./database/sqlite/sqlite-database.db")
	
	if err != nil {
		log.Fatal(err.Error())
	}
	return SqliteDB{
		sqliteDatabase,
	}
}
