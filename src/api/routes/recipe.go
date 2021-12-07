package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	repositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/recipe"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
)

func StartRecipes(router *gin.Engine) {
	RecipeRepository := repositories.NewRecipeRepository(database.InitConnection())

	recipes := router.Group("/api/users/:id/weekmeal")
	{
		recipes.GET("/", handlers.GetRecipes(RecipeRepository))
		// recipes.GET("recipe/:idRecipe", handlers.GetRecipeByID)
		recipes.POST("/recipe", handlers.CreateRecipe(RecipeRepository))
		recipes.PUT("recipe/:idRecipe", handlers.UpdateRecipe(RecipeRepository))
		// recipes.DELETE("recipe/:idRecipe", handlers.DeleteRecipe)
	}
}
