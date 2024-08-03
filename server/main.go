package main

import (
	"app/database"
	"app/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// const db = database.InitializeSqlLite()
	database.InitializeSqlLite()
	router.InitRouter()
}
