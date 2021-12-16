package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	repositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/user"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
)

func StartAuth(router *gin.Engine) {
	userRepository := repositories.NewUserRepository(database.InitConnection())

	auth := router.Group("/auth")
	{
		auth.POST("/signup", handlers.SignUp(userRepository))
		auth.POST("/login", handlers.LogIn(userRepository))
		auth.POST("/logout")
	}
}
