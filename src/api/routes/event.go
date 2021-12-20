package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/fullstacktf/personal-nutritionist-backend/api/middlewares"
	repositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/event"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
)

func StartEvents(router *gin.Engine) {
	EventRepository := repositories.NewEventRepository(database.InitConnection())

	events := router.Group("/calendar", middlewares.IsAuthorized())
	{
		events.GET("/users/:id", handlers.GetEvents(EventRepository))
		events.GET("/event/:idEvent", handlers.GetEventByID(EventRepository))
		events.POST("/event", handlers.CreateEvent(EventRepository))
		events.PUT("/event/:idEvent", handlers.UpdateEvent(EventRepository))
		events.DELETE("/event/:idEvent", handlers.DeleteEvent(EventRepository))
	}
}
