package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
)

var clients = []models.Client{
	{ID: "1", Name: "Train", Dni: "41257854L", Phone: 612732894, Type_diet: "vegan", Weight: 80, Height: 164},
	{ID: "2", Name: "Jeru", Dni: "87654321P", Type_diet: "vegan", Weight: 120, Height: 160},
	{ID: "3", Name: "Sarah Vaughan", Dni: "12345678P", Type_diet: "vegetarian", Weight: 60, Height: 173},
}

func GetClients(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clients)
}

func GetClientbyID(c *gin.Context) {
	id := c.Param("id")

	for _, client := range clients {
		if client.ID == id {
			c.IndentedJSON(http.StatusOK, client)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "client not found 💣"})
}

func PostClients(c *gin.Context) {
	var newClient models.Client

	if err := c.BindJSON(&newClient); err != nil {
		return
	}

	clients = append(clients, newClient)
	c.IndentedJSON(http.StatusCreated, newClient)
}
