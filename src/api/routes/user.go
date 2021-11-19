package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/gin-gonic/gin"
)

func StartUsers(router *gin.Engine) {
	users := router.Group("/api/users")
	{
		users.GET("/", handlers.GetUsers)
		users.GET("/:id", handlers.GetUserByID)
		users.POST("/", handlers.PostUser)
		users.PUT("/:id", handlers.PutUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}
}
