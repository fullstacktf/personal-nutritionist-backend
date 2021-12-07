package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	repositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/recipe"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
)

func StartRecipes(router *gin.Engine) {
	recipeRepository := repositories.NewRecipeRepository(database.InitConnection())

	recipes := router.Group("/api/users/:id/weekmeal")
	{
		recipes.GET("/", handlers.GetRecipes(recipeRepository))
		recipes.GET("recipe/:idRecipe", handlers.GetRecipeByID(recipeRepository))
		recipes.POST("/recipe", handlers.CreateRecipe(recipeRepository))
		recipes.PUT("recipe/:idRecipe", handlers.UpdateRecipe(recipeRepository))
  	recipes.DELETE("recipe/:idRecipe", handlers.DeleteRecipe(recipeRepository))
	}
}
