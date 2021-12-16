package models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	SignUp(c *gin.Context, user *User) (*string, error)
	LogIn(c *gin.Context, credential *Auth) (*string, error)

	GetUsers(c *gin.Context) ([]User, error)
	GetUserByID(c *gin.Context, id primitive.ObjectID) (*User, error)
	GetUsersByRole(c *gin.Context, role string) ([]User, error)
	UpdateUser(c *gin.Context, id primitive.ObjectID, newUser *User) (*User, error)
	DeleteUser(c *gin.Context, id primitive.ObjectID) (*User, error)
}

type EventRepository interface {
	GetEvents(c *gin.Context) ([]Event, error)
	GetEventByID(c *gin.Context, id primitive.ObjectID) (*Event, error)
	CreateEvent(c *gin.Context, event *Event) (primitive.ObjectID, error)
	UpdateEvent(c *gin.Context, id primitive.ObjectID, newEvent *Event) (*Event, error)
	DeleteEvent(c *gin.Context, id primitive.ObjectID) (*Event, error)
}

type RecipeRepository interface {
	GetRecipes(c *gin.Context) ([]Recipe, error)
	GetRecipeByID(c *gin.Context, id primitive.ObjectID) (*Recipe, error)
	CreateRecipe(c *gin.Context, recipe *Recipe) (primitive.ObjectID, error)
	UpdateRecipe(c *gin.Context, id primitive.ObjectID, newRecipe *Recipe) (*Recipe, error)
	DeleteRecipe(c *gin.Context, id primitive.ObjectID) (*Recipe, error)
}
