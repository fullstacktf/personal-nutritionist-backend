package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/gin-gonic/gin"
)

func StartEvents(router *gin.Engine) {
	events := router.Group("/api/users/:id/calendar")
	{
		events.GET("/", handlers.GetEvents)
		events.GET("/event/:idEvent", handlers.GetEventByID)
		events.POST("/event", handlers.PostEvent)
		events.PUT("/event/:idEvent", handlers.PutEvent)
		events.DELETE("/event/:idEvent", handlers.DeleteEvent)
	}
}
