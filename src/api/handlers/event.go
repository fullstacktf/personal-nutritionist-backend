package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
)

var events = []models.Event{
	{ID: "1", Title: "Reunión con el Ghodito", Description: "", Status: "Confirmado", StartingDate: "8-8-2021 13:00", EndingDate: "8-8-2021 13:30"},
	{ID: "2", Title: "Seguimiento de dieta", Description: "Siguiendo avances de lola", Status: "Confirmado", StartingDate: "41257854L", EndingDate: "612732894"},
}

func GetEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

func GetEventByID(c *gin.Context) {
	id := c.Param("idEvent")

	for _, event := range events {
		if event.ID == id {
			c.IndentedJSON(http.StatusOK, event)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found 💣"})
}

func PostEvent(c *gin.Context) {
	var newEvent models.Event
	if err := c.BindJSON(&newEvent); err != nil {
		return
	}

	events = append(events, newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func PutEvent(c *gin.Context) {
	id := c.Param("idEvent")

	var updatedEvent models.Event
	if err := c.BindJSON(&updatedEvent); err != nil {
		return
	}

	for index, event := range events {
		if id == event.ID {
			events[index] = updatedEvent
			c.IndentedJSON(http.StatusOK, updatedEvent)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found 💣"})
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("idEvent")

	for index, event := range events {
		if id == event.ID {
			events = append(events[:index], events[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "event deleted 🧹"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found 💣"})
}
