package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: "1", Name: "Sergio Peinado", Email: "sergiopeinado@gmail.com", Role: "Nutritionist", Username: "Sergito", Dni: "41257854L", Phone: 612732894, Likes: 157770, Is_verified: true},
	{ID: "2", Name: "Godhito", Email: "damecomidah@gmail.com", Dni: "87654321P", Username: "Adanito", Type_diet: "Hypercaloric", Weight: 120, Height: 160, Role: "Client"},
	{ID: "3", Name: "Sarah Vaughan", Dni: "12345678P", Type_diet: "vegetarian", Weight: 60, Height: 173, Role: "Client"},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func GetUserbyID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found 💣"})
}

func PostUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func PutUser(c *gin.Context) {
	id := c.Param("id")

	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		return
	}

	for index, user := range users {
		if id == user.ID {
			users[index] = updatedUser
			c.IndentedJSON(http.StatusOK, updatedUser)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found 💣"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for index, user := range users {
		if id == user.ID {
			users = append(users[:index], users[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted 🧹"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found 💣"})
}
