package services

import (
	"errors"
	"log"
	"time"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/env"
	"github.com/golang-jwt/jwt/v4"
)

type AuthClaims struct {
	Authorized bool   `json:"authorized"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(email, role string) (*models.Token, error) {
	key := []byte(env.JWT_SECRET)
	claims := AuthClaims{
		true,
		email,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(time.Minute*30).Unix(), 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}
	tokenStruct := models.Token{Email: email, Role: role, TokenString: signedToken}

	return &tokenStruct, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	key := []byte(env.JWT_SECRET)

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("invalid token")
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	log.Println(("Llego aqui"))

	if claims, isValid := token.Claims.(jwt.MapClaims); isValid && token.Valid {
		if claims["role"] == "Nutricionista" {
			log.Panicln("Eres un nutricionista")
		} else if claims["role"] == "Cliente" {
			log.Panicln("Eres un cliente")
		}
	}

	return token, nil
}
