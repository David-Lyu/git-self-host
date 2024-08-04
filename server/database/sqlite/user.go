package sqliteDriver

import (
	// "app/database"
	"database/sql"
	"log"
	
	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (sd SqliteDB) createUserTable() {
	createUserTableSQL := `CREATE TABLE User (
		"id" TEXT UNIQUE NOT NULL PRIMARY KEY,
		"email" TEXT,
		"username" TEXT,
		"password" TEXT,
		"isActive" BOOLEAN
	  );`
	
	log.Println("Create user table...")
	statement, err := sd.db.Prepare(createUserTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	if _, err := statement.Exec(); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("user table created")
}

func (sd SqliteDB) InsertUser(c *gin.Context) {
	id := uuid.New()
	
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	// /* TODO  Password Hashing */
	insertUserSQL := `INSERT INTO User(id, email, username, password, isActive) VALUES (?, ?, ?, ?, ?)`
	statement, err := sd.db.Prepare(insertUserSQL)
	
	// statement, err := database.DB.InsertDB("User", database.DB.Db)
	
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, user.Email, user.Username, user.Password, true)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.Fatalln(err.Error())
	}
	
	c.JSON(200, gin.H{"message": "User inserted!"})
}

func (sd SqliteDB) GetUsers(c *gin.Context) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./database/sqlite/sqlite-database.db")
	defer sqliteDatabase.Close()
	
	row, err := sqliteDatabase.Query("SELECT * FROM User")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.Fatal(err)
	}
	defer row.Close()
	
	users := make([]*models.User, 0)
	
	for row.Next() {
		user := new(models.User)
		err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.IsActive)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			log.Fatal(err)
		}
		users = append(users, user)
	}
	
	c.JSON(200, users)
}

func (sd SqliteDB) GetUser(c *gin.Context) {
	id := c.Param("id")
	
	sqliteDatabase, _ := sql.Open("sqlite3", "./database/sqlite/sqlite-database.db")
	defer sqliteDatabase.Close()
	
	row, err := sqliteDatabase.Query("SELECT * FROM User where id = ? ", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.Fatal(err)
	}
	defer row.Close()
	
	users := make([]*models.User, 0)
	
	for row.Next() {
		user := new(models.User)
		err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.IsActive)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			log.Fatal(err)
		}
		users = append(users, user)
	}
	
	c.JSON(200, users)
}

func (sd SqliteDB) UpdateUser(c *gin.Context) {
	/* TODO Implement the update functionality */
	
	c.JSON(200, gin.H{"message": "update user"})
}

func (sd SqliteDB) DeleteUser(c *gin.Context) {
	/* TODO Implement the delete functionality */
	
	c.JSON(200, gin.H{"message": "delete user"})
}
