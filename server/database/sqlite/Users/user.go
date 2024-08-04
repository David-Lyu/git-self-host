package sqliteUsers

import (
	// "app/database"
	"database/sql"
)

func InsertUsers(database *sql.DB) (*sql.Stmt, error) {
	insertUserSQL := `INSERT INTO User(id, email, username, password, isActive) VALUES (?, ?, ?, ?, ?)`
	return database.Prepare(insertUserSQL)
}
