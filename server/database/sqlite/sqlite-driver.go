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

func createTable(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE User (
		"id" TEXT UNIQUE NOT NULL PRIMARY KEY,
		"email" TEXT,
		"username" TEXT,
		"password" TEXT,
		"isActive" BOOLEAN
	  );` // SQL Statement for Create Table
	
	log.Println("Create user table...")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("user table created")
}

func NewSqliteDB() SqliteDB {
	os.Remove("./database/sqlite/sqlite-database.db")
	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("./database/sqlite/sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")
	
	sqliteDatabase, _ := sql.Open("sqlite3", "./database/sqlite/sqlite-database.db")
	// defer sqliteDatabase.Close()
	createTable(sqliteDatabase)
	
	return SqliteDB{
		sqliteDatabase,
	}
}
