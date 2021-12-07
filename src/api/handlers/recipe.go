package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRecipes(repository models.RecipeRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		recipes, err := repository.GetRecipes(c)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, recipes)
		}
	}
}

// func GetRecipes(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, recipes)
// }

// func GetRecipeByID(c *gin.Context) {
// 	id := c.Param("idRecipe")

// 	for _, recipe := range recipes {
// 		if recipe.ID == id {
// 			c.IndentedJSON(http.StatusOK, recipe)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found 💣"})
// }

func CreateRecipe(repository models.RecipeRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var recipe models.Recipe
		if err := c.BindJSON(&recipe); err != nil {
			return
		}

		objectId, err := repository.CreateRecipe(c, &recipe)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, objectId)
		}
	}
}

// func PutRecipe(c *gin.Context) {
// 	id := c.Param("idRecipe")

// 	var updatedRecipe models.Recipe
// 	if err := c.BindJSON(&updatedRecipe); err != nil {
// 		return
// 	}

// 	for index, recipe := range recipes {
// 		if id == recipe.ID {
// 			recipes[index] = updatedRecipe
// 			c.IndentedJSON(http.StatusOK, updatedRecipe)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found 💣"})
// }

func DeleteRecipe(repository models.RecipeRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("idRecipe"))

		event, err := repository.DeleteRecipe(c, id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, event)
		}
	}
}
