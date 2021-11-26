package models

import "github.com/gin-gonic/gin"

type BasicUser struct {
	ID         string `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Phone      uint32 `json:"phone"`
	Photo      string `json:"photo"`
	IsVerified bool   `json:"isVerified"`
}

type User struct {
	ID          string   `json:"id" binding:"required"`
	Email       string   `json:"email" binding:"required"`
	Username    string   `json:"username" binding:"required"`
	Password    string   `json:"password" binding:"required"`
	Role        string   `json:"role" binding:"required"`
	Photo       string   `json:"photo"`
	Name        string   `json:"name" binding:"required"`
	Dni         string   `json:"dni"`
	Birthday    string   `json:"birthday"`
	Phone       uint32   `json:"phone"`
	Description string   `json:"description"`
	Events      []Event  `json:"events"`
	Recipes     []Recipe `json:"recipes"`

	// Nutricionist
	IsVerified  bool        `json:"isVerified"`
	Education   []string    `json:"education"`
	Specialties []string    `json:"specialties"`
	Price       float64     `json:"price"`
	Likes       uint64      `json:"likes"`
	Clients     []BasicUser `json:"clients"`

	// Client
	Weight        uint        `json:"weight"`
	Height        uint        `json:"height"`
	TypeDiet      string      `json:"typeDiet"`
	Intolerances  []string    `json:"intolerances"`
	Nutricionists []BasicUser `json:"nutricionists"`
}

type UserRepository interface {
	GetUsers(context *gin.Context) []User
}
