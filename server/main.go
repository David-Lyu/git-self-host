package main

import (
	"app/database"
	"app/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitializeSqlLite()
	router.InitRouter()
}
