package models

import "github.com/gin-gonic/gin"

type CollectionDB interface {
	InsertCollection(c *gin.Context)
	GetCollections(c *gin.Context)
	GetCollection(c *gin.Context)
	UpdateCollection(c *gin.Context)
	DeleteCollection(c *gin.Context)
}
