package handlers

import (
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(repository models.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			return
		}

		msg, err := repository.SignUp(c, &user)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, msg)
		}
	}
}

func LogIn(repository models.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credential models.Auth
		if err := c.BindJSON(&credential); err != nil {
			return
		}

		msg, err := repository.LogIn(c, &credential)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, msg)
		}
	}
}

func GetUsers(repository models.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := repository.GetUsers(c)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, users)
		}
	}
}

func GetUserByID(repository models.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		user, err := repository.GetUserByID(c, id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	}
}

func GetUsersByRole(repository models.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.Param("role")

		users, err := repository.GetUsersByRole(c, role)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, users)
		}
	}
}

func UpdateUser(repository models.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		var newUser models.User
		if err := c.BindJSON(&newUser); err != nil {
			return
		}

		user, err := repository.UpdateUser(c, id, &newUser)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	}
}

func DeleteUser(repository models.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		user, err := repository.DeleteUser(c, id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "💣", "message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	}
}
