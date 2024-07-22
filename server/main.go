package main

import (
	"database/sql"
	"log"
	"os"

	"app/models"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	initializeSqlLite()

	r := gin.Default()

	router := r.Group("/")
	{
		router.POST("/user/", insertUser)
		router.GET("/getUsers/", getUsers)
	}

	r.Run()

}

func insertUser(c *gin.Context) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./models/database/sqlite-database.db")
	defer sqliteDatabase.Close()

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	insertUserSQL := `INSERT INTO User(email, username, isActive) VALUES (?, ?, ?)`
	statement, err := sqliteDatabase.Prepare(insertUserSQL)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(user.Email, user.Username, true)
	if err != nil {
		log.Fatalln(err.Error())
	}

	c.JSON(200, gin.H{"message": "User inserted!"})
}

func getUsers(c *gin.Context) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./models/database/sqlite-database.db")
	defer sqliteDatabase.Close()

	row, err := sqliteDatabase.Query("SELECT * FROM User")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	users := make([]*models.User, 0)

	for row.Next() {
		user := new(models.User)
		err := row.Scan(&user.Id, &user.Email, &user.Username, &user.IsActive)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	c.JSON(200, gin.H{"users": users})
}

func initializeSqlLite() {
	os.Remove("./models/database/sqlite-database.db")
	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("./models/database/sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./models/database/sqlite-database.db")
	defer sqliteDatabase.Close()
	createTable(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE User (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
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
