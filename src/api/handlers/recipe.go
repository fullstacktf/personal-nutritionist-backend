package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
)

var recipes = []models.Recipe{
	{ID: "1", Name: "Porridge de nueces y arandanos", TypeDiet: "Vegana", TypeMeal: "Desayuno", Alergens: []string{"Frutos secos"}, Date: "24-12-2021", Ingredients: []string{"Nueces", "Arándanos"}},
	{ID: "2", Name: "Papas rellenas", TypeMeal: "Almuerzo", Alergens: []string{"Carne de cerdo", "Carne de vaca"}, Date: "12-12-2021", Ingredients: []string{"Papas", "Carne", "Mojo"}},
}

func GetRecipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, recipes)
}

func GetRecipeByID(c *gin.Context) {
	id := c.Param("idRecipe")

	for _, recipe := range recipes {
		if recipe.ID == id {
			c.IndentedJSON(http.StatusOK, recipe)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found 💣"})
}

func PostRecipe(c *gin.Context) {
	var newRecipe models.Recipe
	if err := c.BindJSON(&newRecipe); err != nil {
		return
	}

	recipes = append(recipes, newRecipe)
	c.IndentedJSON(http.StatusCreated, newRecipe)
}

func PutRecipe(c *gin.Context) {
	id := c.Param("idRecipe")

	var updatedRecipe models.Recipe
	if err := c.BindJSON(&updatedRecipe); err != nil {
		return
	}

	for index, recipe := range recipes {
		if id == recipe.ID {
			recipes[index] = updatedRecipe
			c.IndentedJSON(http.StatusOK, updatedRecipe)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found 💣"})
}

func DeleteRecipe(c *gin.Context) {
	id := c.Param("idRecipe")

	for index, recipe := range recipes {
		if id == recipe.ID {
			recipes = append(recipes[:index], recipes[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "recipe deleted 🧹"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found 💣"})
}
