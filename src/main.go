package main

import (
	"fmt"
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api", func(c *gin.Context) {
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
	err := router.Run(":8080")

	if err != nil {
		fmt.Println("Funciono")
	}
}
