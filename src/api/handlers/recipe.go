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

func UpdateRecipe(repository models.RecipeRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("idRecipe"))
		var newRecipe models.Recipe
		if err := c.BindJSON(&newRecipe); err != nil {
			return
		}

		recipe, err := repository.UpdateRecipe(c, id, &newRecipe)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, recipe)
		}
	}
}

// func DeleteRecipe(c *gin.Context) {
// 	id := c.Param("idRecipe")

// 	for index, recipe := range recipes {
// 		if id == recipe.ID {
// 			recipes = append(recipes[:index], recipes[index+1:]...)
// 			c.IndentedJSON(http.StatusOK, gin.H{"message": "recipe deleted 🧹"})
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found 💣"})
// }
