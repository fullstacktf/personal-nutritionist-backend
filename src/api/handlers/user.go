package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/api/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(repository repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := repository.GetUsers()
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, users)
		}
	}
}

func GetUserByID(repository repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		user, err := repository.GetUserByID(id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	}
}

func PostUser(repository repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			return
		}

		objectId, err := repository.PostUser(&user)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, objectId)
		}
	}
}

func PutUser(repository repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		var newUser models.User
		if err := c.BindJSON(&newUser); err != nil {
			return
		}

		user, err := repository.PutUser(id, newUser)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	}
}

func DeleteUser(repository repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		user, err := repository.DeleteUser(id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	}
}
