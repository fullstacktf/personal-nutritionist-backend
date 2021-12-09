package main

import (
	"log"
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Nutriguide mi cuaaaate 🌮🤠🥑",
		})
	})

	routes.StartUsers(router)
	routes.StartEvents(router)
	routes.StartRecipes(router)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	err := router.Run(":5000")
	if err != nil {
		log.Fatalln("Error running on port 💣:", err)
	}
}
