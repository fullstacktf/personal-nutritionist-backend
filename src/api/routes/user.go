package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	// "github.com/fullstacktf/personal-nutritionist-backend/api/middlewares"
	repositoriesUser "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/user"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
)

func StartUsers(router *gin.Engine) {
	userRepository := repositoriesUser.NewUserRepository(database.InitConnection())

	// users := router.Group("/users", middlewares.IsAuthorized(authRepository))
	users := router.Group("/users")
	{
		users.GET("/", handlers.GetUsers(userRepository))
		users.GET("/:id", handlers.GetUserByID(userRepository))
		users.GET("/role/:role", handlers.GetUsersByRole(userRepository))
		users.PUT("/:id", handlers.UpdateUser(userRepository))
		users.DELETE("/:id", handlers.DeleteUser(userRepository))
	}
}
