package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	repositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/event"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
)

func StartEvents(router *gin.Engine) {
	EventRepository := repositories.NewEventRepository(database.InitConnection())

	events := router.Group("/api/users/:id/calendar")
	{
		events.GET("/", handlers.GetEvents(EventRepository))
		// events.GET("/event/:idEvent", handlers.GetEventByID)
		// events.POST("/event", handlers.PostEvent)
		// events.PUT("/event/:idEvent", handlers.PutEvent)
		// events.DELETE("/event/:idEvent", handlers.DeleteEvent)
	}
}
