package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	repositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/user"
	"github.com/gin-gonic/gin"
)

func StartUsers(router *gin.Engine) {
	userRepository := repositories.NewUserRepository(nil)

	users := router.Group("/api/users")
	{
		users.GET("/", handlers.GetUsers(userRepository))
		users.GET("/:id", handlers.GetUserByID(userRepository))
		users.POST("/", handlers.PostUser(userRepository))
		users.PUT("/:id", handlers.PutUser(userRepository))
		users.DELETE("/:id", handlers.DeleteUser(userRepository))
	}
}
