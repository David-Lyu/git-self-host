package models

import "github.com/gin-gonic/gin"

type Repo struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	IsPrivate bool   `json:"is_private"`
}

type RepoDB interface {
	InsertRepo(c *gin.Context)
	GetRepos(c *gin.Context)
	GetRepo(c *gin.Context)
	UpdateRepo(c *gin.Context)
	DeleteRepo(c *gin.Context)
}
