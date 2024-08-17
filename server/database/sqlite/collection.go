package sqliteDriver

import (
	// "app/database"

	"log"

	"app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
*
Create Collection with user ID,
Read Collection (get/ gets)
Update Collection
Delete Collection
*/
func (sd SqliteDB) createCollection(c *gin.Context) {
	var id = uuid.UUID.NodeID("test")
	sd.db.Exec("INSERT ", "")
	log.Default("Created Collection")
}

func (sd SqliteDB) getCollect(c *gin.Context) {
	// var user, _ = c.Params.Get("user");
	var loginKey, _ = c.Param("isLoggedIn")
	//grab USER
	var user, _ = sd.GetUser(c * gin.Context)
	if loginKey == models.User.IsActive {
		println("same same")
	}

}
