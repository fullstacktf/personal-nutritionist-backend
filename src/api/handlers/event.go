package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
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
