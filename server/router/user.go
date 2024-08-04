package router

import (
	"github.com/gin-gonic/gin"
)

func (r *Router) addUserRoute(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	
	userRoute.POST("/", r.db.InsertUser)
	userRoute.GET("/", r.db.GetUsers)
	userRoute.GET("/:id", r.db.GetUser)
	userRoute.PATCH("/:id", r.db.UpdateUser)
	userRoute.DELETE("/:id", r.db.DeleteUser)
	
}
