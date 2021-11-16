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

func UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var UpdatedClient models.Client

	if err := c.BindJSON(&UpdatedClient); err != nil {
		return
	}

	for index, client := range clients {
		if id == client.ID {
			clients[index] = UpdatedClient
			c.IndentedJSON(http.StatusOK, UpdatedClient)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "client not found 💣"})
}

func DeleteClient(c *gin.Context) {
	id := c.Param("id")
	var DeletedClient models.Client
	for index, client := range clients {
		if id == client.ID {
			copy(clients[index:], clients[index+1:])
			clients[len(clients)-1] = DeletedClient
			clients = clients[:len(clients)-1]
			c.IndentedJSON(http.StatusOK, gin.H{"message": "client deleted 🧹"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "client not found 💣"})
}
