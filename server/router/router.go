package router

import (
	"app/database"
	"github.com/gin-gonic/gin"
)

type Router struct {
	routes *gin.Engine
	db     database.DB
}

func NewRouter(db database.DB) Router {
	r := Router{
		routes: gin.Default(),
		db:     db,
	}
	
	v1 := r.routes.Group("/v1")
	
	r.addUserRoute(v1)
	r.addRepoRoute(v1)
	
	return r
}

func (r *Router) Run() error {
	return r.routes.Run()
}
