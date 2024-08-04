package main

import (
	"log"
	
	"app/database"
	"app/router"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	h := router.NewRouter(database.InitDatabase())
	if err := h.Run(); err != nil {
		log.Fatal(err)
	}
}
