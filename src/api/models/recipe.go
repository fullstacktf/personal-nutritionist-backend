package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
	ObjectID    primitive.ObjectID `json:"_id" bson:"_id"`
	Owner       primitive.ObjectID `json:"owner" bson:"owner" binding:"required"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Date        string             `json:"date" bson:"date"`
	TypeMeal    string             `json:"typeMeal" bson:"typeMeal"`
	TypeDiet    string             `json:"typeDiet" bson:"typeDiet"`
	Alergens    []string           `json:"alergens" bson:"alergens" binding:"required"`
	Ingredients []string           `json:"ingredients" bson:"ingredients" binding:"required"`
	Preparation string             `json:"preparation" bson:"preparation"`
	CookingTime string             `json:"cookingTime" bson:"cookingTime"`
}
