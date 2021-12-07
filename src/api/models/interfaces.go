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

type EventRepository interface {
	GetEvents(c *gin.Context) ([]Event, error)
	GetEventByID(c *gin.Context, id primitive.ObjectID) (*Event, error)
	CreateEvent(c *gin.Context, event *Event) (primitive.ObjectID, error)
	UpdateEvent(c *gin.Context, id primitive.ObjectID, newEvent *Event) (*Event, error)
}

type RecipeRepository interface {
	CreateRecipe(c *gin.Context, recipe *Recipe) (primitive.ObjectID, error)
}
