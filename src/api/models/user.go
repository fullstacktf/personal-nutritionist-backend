package models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BasicUser struct {
	ObjectID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	Email      string             `json:"email" bson:"email" binding:"required"`
	Phone      uint32             `json:"phone" bson:"phone"`
	Photo      string             `json:"photo" bson:"photo"`
	IsVerified bool               `json:"isVerified" bson:"isVerified"`
}

type User struct {
	ObjectID    primitive.ObjectID `json:"_id" bson:"_id"`
	Email       string             `json:"email" bson:"email" binding:"required"`
	Username    string             `json:"username" bson:"username" binding:"required"`
	Password    string             `json:"password" bson:"password" binding:"required"`
	Role        string             `json:"role" bson:"role" binding:"required"`
	Photo       string             `json:"photo" bson:"photo"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Dni         string             `json:"dni" bson:"dni"`
	Birthday    string             `json:"birthday" bson:"birthday"`
	Phone       uint32             `json:"phone" bson:"phone"`
	Description string             `json:"description" bson:"description"`
	Events      []Event            `json:"events" bson:"events"`
	Recipes     []Recipe           `json:"recipes" bson:"recipes"`

	// Nutricionist
	IsVerified  bool        `json:"isVerified" bson:"isVerified"`
	Education   []string    `json:"education" bson:"education"`
	Specialties []string    `json:"specialties" bson:"specialties"`
	Price       float64     `json:"price" bson:"price"`
	Likes       uint64      `json:"likes" bson:"likes"`
	Clients     []BasicUser `json:"clients" bson:"clients"`

	// Client
	Weight        uint        `json:"weight" bson:"weight"`
	Height        uint        `json:"height" bson:"height"`
	TypeDiet      string      `json:"typeDiet" bson:"typeDiet"`
	Intolerances  []string    `json:"intolerances" bson:"intolerances"`
	Nutricionists []BasicUser `json:"nutricionists" bson:"nutricionists"`
}

type UserRepository interface {
	GetUsers(c *gin.Context) ([]User, error)
	GetUserByID(c *gin.Context, id primitive.ObjectID) (User, error)
	PostUser(c *gin.Context, user *User) (primitive.ObjectID, error)
	PutUser(c *gin.Context, id primitive.ObjectID, newUser User) (User, error)
	DeleteUser(c *gin.Context, id primitive.ObjectID) (User, error)
}
