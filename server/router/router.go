package router

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() {

	r := gin.Default()

	router := r.Group("/v1")
	{
		// users
		user_router := router.Group("/user")
		{
			user_router.POST("/", controllers.InsertUser)
			user_router.GET("/", controllers.GetUsers)
			user_router.GET("/:id", controllers.GetUser)
			user_router.PATCH("/", controllers.UpdateUser)
			user_router.DELETE("/", controllers.DeleteUser)
		}
	}

	r.Run()
}
