package controllers

import (
	"database/sql"
	"log"

	"app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func InsertUser(c *gin.Context) {
	id := uuid.New()

	sqliteDatabase, _ := sql.Open("sqlite3", "./database/sqlite-database.db")
	defer sqliteDatabase.Close()

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	/* TODO  Password Hashing */
	insertUserSQL := `INSERT INTO User(id, email, username, password, isActive) VALUES (?, ?, ?, ?, ?)`
	statement, err := sqliteDatabase.Prepare(insertUserSQL)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, user.Email, user.Username, true)
	if err != nil {
		log.Fatalln(err.Error())
	}

	c.JSON(200, gin.H{"message": "User inserted!"})
}

func GetUsers(c *gin.Context) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./database/sqlite-database.db")
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

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	sqliteDatabase, _ := sql.Open("sqlite3", "./database/sqlite-database.db")
	defer sqliteDatabase.Close()

	row, err := sqliteDatabase.Query("SELECT * FROM User where id = ? ", id)
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

	c.JSON(200, users)
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "update user"})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "delete user"})
}
