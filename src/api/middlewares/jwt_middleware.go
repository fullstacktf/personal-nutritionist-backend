package middlewares

import (
	"fmt"
	"net/http"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func IsAuthorized(repository models.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "💣", "message": "no token found"})
		}

		token, err := services.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "💣", "message": err})
		}
		// token, err := services.ValidateToken(tokenString)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "💣", "message": err})
		// }

		// c.IndentedJSON(http.StatusOK, token)
	}
}
