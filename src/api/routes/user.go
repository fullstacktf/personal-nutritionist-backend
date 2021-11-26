package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/fullstacktf/personal-nutritionist-backend/api/repositories"
	"github.com/gin-gonic/gin"
)

func StartUsers(router *gin.Engine) {
	//	db, err := sql.Open("mysql", mysqlURI)
	userRepository := repositories.NewUserRepository(nil) //TODO replace nil to db

	users := router.Group("/api/users")
	{
		users.GET("/", handlers.GetUsers(userRepository))
		users.GET("/:id", handlers.GetUserByID)
		users.POST("/", handlers.PostUser)
		users.PUT("/:id", handlers.PutUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}
}
