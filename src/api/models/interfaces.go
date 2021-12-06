package models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	GetUsers(c *gin.Context) ([]User, error)
	GetUserByID(c *gin.Context, id primitive.ObjectID) (*User, error)
	CreateUser(c *gin.Context, user *User) (primitive.ObjectID, error)
	UpdateUser(c *gin.Context, id primitive.ObjectID, newUser *User) (*User, error)
	DeleteUser(c *gin.Context, id primitive.ObjectID) (*User, error)
}
