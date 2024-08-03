package sqliteDriver

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func sqliteDriver() {

	// os.Remove("./database/sqlite-database.db")
	// log.Println("Creating sqlite-database.db...")
	// file, err := os.Create("./database/sqlite-database.db") // Create SQLite file
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// file.Close()
	// log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./database/sqlite-database.db")
	defer sqliteDatabase.Close()
	createTable(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE User (
		"id" TEXT UNIQUE NOT NULL PRIMARY KEY,
		"email" TEXT,
		"username" TEXT,
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

func closeDB(sql *sql.DB) {
	sql.close()
}
