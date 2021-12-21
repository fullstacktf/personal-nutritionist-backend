package routes

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/fullstacktf/personal-nutritionist-backend/api/middlewares"
	repositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/event"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
)

func StartEvents(router *gin.Engine) {
	eventRepository := repositories.NewEventRepository(database.InitConnection())

	events := router.Group("/calendar", middlewares.IsAuthorized())
	{
		events.GET("/users/:id", handlers.GetEvents(eventRepository))
		events.GET("/event/:idEvent", handlers.GetEventByID(eventRepository))
		events.POST("/event", handlers.CreateEvent(eventRepository))
		events.PUT("/event/:idEvent", handlers.UpdateEvent(eventRepository))
		events.DELETE("/event/:idEvent", handlers.DeleteEvent(eventRepository))
	}
}
