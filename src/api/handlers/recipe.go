package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/services"
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

func GetRecipeByID(repository models.RecipeRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("idRecipe"))

		user, err := repository.GetRecipeByID(c, id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	}
}

func CreateRecipe(repository models.RecipeRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newRecipe models.Recipe
		if err := c.BindJSON(&newRecipe); err != nil {
			return
		}

		valid := services.ValidateData(newRecipe)
		if !valid {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": "invalid data inputs"})
			return
		}

		recipe, err := repository.CreateRecipe(c, &newRecipe)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, recipe)
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

		valid := services.ValidateData(newRecipe)
		if !valid {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": "invalid data inputs"})
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
