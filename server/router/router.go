package router

import (
	"app/database"
	"github.com/gin-gonic/gin"
)

type Router struct {
	routes *gin.Engine
	db     database.Databaser
}

func NewRouter(db database.Databaser) Router {
	r := Router{
		routes: gin.Default(),
		db:     db,
	}
	
	v1 := r.routes.Group("/v1")
	
	r.addUserRoute(v1)
	
	return r
}

func (r *Router) Run() error {
	return r.routes.Run()
}
