package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/gin-gonic/gin"
)

func StartRecipes(router *gin.Engine) {
	recipes := router.Group("/api/users/:id/weekmeal")
	{
		recipes.GET("/", handlers.GetRecipes)
		recipes.GET("recipe/:idRecipe", handlers.GetRecipeByID)
		recipes.POST("recipe", handlers.PostRecipe)
		recipes.PUT("recipe/:idRecipe", handlers.PutRecipe)
		recipes.DELETE("recipe/:idRecipe", handlers.DeleteRecipe)
	}
}
