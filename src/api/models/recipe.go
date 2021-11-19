package models

type Recipe struct {
	ID          string   `json:"id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Date        string   `json:"date" binding:"required"`
	TypeMeal    string   `json:"typeMeal" binding:"required"`
	TypeDiet    string   `json:"typeDiet"`
	Alergens    []string `json:"alergens" binding:"required"`
	Ingredients []string `json:"ingredients" binding:"required"`
	Preparation string   `json:"preparation"`
	CookingTime string   `json:"cookingTime"`
}
