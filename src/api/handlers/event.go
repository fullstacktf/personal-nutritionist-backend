package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/services"
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
		var newEvent models.Event
		if err := c.BindJSON(&newEvent); err != nil {
			return
		}

		valid := services.ValidateData(newEvent)
		if !valid {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": "invalid data inputs"})
			return
		}

		event, err := repository.CreateEvent(c, &newEvent)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, event)
		}
	}
}

func UpdateEvent(repository models.EventRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("idEvent"))
		var newEvent models.Event
		if err := c.BindJSON(&newEvent); err != nil {
			return
		}

		valid := services.ValidateData(newEvent)
		if !valid {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": "invalid data inputs"})
			return
		}

		event, err := repository.UpdateEvent(c, id, &newEvent)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, event)
		}
	}
}

func DeleteEvent(repository models.EventRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("idEvent"))

		event, err := repository.DeleteEvent(c, id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, event)
		}
	}
}
