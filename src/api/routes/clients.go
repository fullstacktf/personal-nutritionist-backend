package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/gin-gonic/gin"
)

func StartClients(router *gin.Engine) {
	clients := router.Group("/api/clients")
	{
		clients.GET("/", handlers.GetClients)
		clients.GET("/:id", handlers.GetClientbyID)
		clients.POST("/", handlers.PostClient)
		clients.PUT("/:id", handlers.UpdateClient)
		clients.DELETE("/:id", handlers.DeleteClient)
	}
}
