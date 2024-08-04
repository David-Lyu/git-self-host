package models

import "github.com/gin-gonic/gin"

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
	IsActive bool   `json:"isActive"`
}

type UserDB interface {
	InsertUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
