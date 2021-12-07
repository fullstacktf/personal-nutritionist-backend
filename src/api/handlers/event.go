package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetEvents(repository models.EventRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		events, err := repository.GetEvents(c)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, events)
		}
	}
}

func GetEventByID(repository models.EventRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("idEvent"))

		event, err := repository.GetEventByID(c, id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, event)
    }
  }
}
      
func CreateEvent(repository models.EventRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var event models.Event
		if err := c.BindJSON(&event); err != nil {
			return
		}

		objectId, err := repository.CreateEvent(c, &event)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, objectId)
		}
	}
}
