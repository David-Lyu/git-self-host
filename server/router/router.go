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
		userRouter := router.Group("/user")
		{
			userRouter.POST("/", controllers.InsertUser)
			userRouter.GET("/", controllers.GetUsers)
			userRouter.GET("/:id", controllers.GetUser)
			userRouter.PATCH("/:id", controllers.UpdateUser)
			userRouter.DELETE("/:id", controllers.DeleteUser)
		}
	}

	r.Run()
}
