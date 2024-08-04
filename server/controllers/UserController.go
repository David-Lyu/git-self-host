package controllers

import (
	"database/sql"
	"log"

	"app/database"
	"app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func InsertUser(c *gin.Context) {
	id := uuid.New()

	// sqliteDatabase, _ := sql.Open("sqlite3", "./database/sqlite-database.db")
	// database.DB.db.Open("")
	// defer sqliteDatabase.Close()

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// /* TODO  Password Hashing */
	// insertUserSQL := `INSERT INTO User(id, email, username, password, isActive) VALUES (?, ?, ?, ?, ?)`
	// statement, err := sqliteDatabase.Prepare(insertUserSQL)
	log.Print(database.DB)
	statement, err := database.DB.InsertDB("User", database.DB.Db)

	if err != nil {
		log.Fatalln(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	_, err = statement.Exec(id, user.Email, user.Username, user.Password, true)
	if err != nil {
		log.Fatalln(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"message": "User inserted!"})
}

func GetUsers(c *gin.Context) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./database/sqlite/sqlite-database.db")
	defer sqliteDatabase.Close()

	row, err := sqliteDatabase.Query("SELECT * FROM User")
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer row.Close()

	users := make([]*models.User, 0)

	for row.Next() {
		user := new(models.User)
		err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.IsActive)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"error": err.Error()})
		}
		users = append(users, user)
	}

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	sqliteDatabase, _ := sql.Open("sqlite3", "./database/sqlite/sqlite-database.db")
	defer sqliteDatabase.Close()

	row, err := sqliteDatabase.Query("SELECT * FROM User where id = ? ", id)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer row.Close()

	users := make([]*models.User, 0)

	for row.Next() {
		user := new(models.User)
		err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.IsActive)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"error": err.Error()})
		}
		users = append(users, user)
	}

	c.JSON(200, users)
}

func UpdateUser(c *gin.Context) {
	/* TODO Implement the update functonanlity */

	c.JSON(200, gin.H{"message": "update user"})
}

func DeleteUser(c *gin.Context) {
	/* TODO Implement the delete functonanlity */

	c.JSON(200, gin.H{"message": "delete user"})
}
