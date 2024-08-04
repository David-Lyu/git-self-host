package router

import "github.com/gin-gonic/gin"

func (r *Router) addRepoRoute(rg *gin.RouterGroup) {
	repoRoute := rg.Group("/repo")
	
	repoRoute.POST("/:userId", r.db.InsertRepo)
	repoRoute.GET("/:userId", r.db.GetRepos)
	repoRoute.GET("/:userId/:id", r.db.GetRepo)
	repoRoute.PATCH("/:userId/:id", r.db.UpdateRepo)
	repoRoute.DELETE("/:userId/:id", r.db.DeleteRepo)
}
