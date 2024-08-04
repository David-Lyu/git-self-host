package sqliteDriver

import (
	"github.com/gin-gonic/gin"
)

func (sd SqliteDB) createRepoTable() {
	/* TODO Implement the update functionality */
	/*
		createRepoTableSQL := `CREATE TABLE Repo (
		  );`
	
		log.Println("Create repo table...")
		statement, err := sd.db.Prepare(createRepoTableSQL)
		if err != nil {
			log.Fatal(err.Error())
		}
	
		if _, err := statement.Exec(); err != nil {
			log.Fatal(err.Error())
		}
		log.Println("repo table created")
	
	
	*/
}

func (sd SqliteDB) InsertRepo(c *gin.Context) {
	/* TODO Implement the update functionality */
	
	c.JSON(200, gin.H{"message": "insert repo"})
}

func (sd SqliteDB) GetRepos(c *gin.Context) {
	/* TODO Implement the update functionality */
	
	c.JSON(200, gin.H{"message": "get repos"})
}

func (sd SqliteDB) GetRepo(c *gin.Context) {
	/* TODO Implement the update functionality */
	
	c.JSON(200, gin.H{"message": "get repo"})
}

func (sd SqliteDB) UpdateRepo(c *gin.Context) {
	/* TODO Implement the update functionality */
	
	c.JSON(200, gin.H{"message": "update repo"})
}

func (sd SqliteDB) DeleteRepo(c *gin.Context) {
	/* TODO Implement the delete functionality */
	
	c.JSON(200, gin.H{"message": "delete repo"})
}
