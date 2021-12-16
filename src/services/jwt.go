package services

import (
	"errors"
	"log"
	"time"

	"github.com/fullstacktf/personal-nutritionist-backend/env"
	"github.com/golang-jwt/jwt/v4"
)

type AuthClaims struct {
	Authorized bool   `json:"authorized"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(email, role string) (*string, error) {
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

	return &signedToken, nil
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

// // func parseToken() {
// // 	receivedToken := "xxxxx.yyyyy.zzzzz"

// // 	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
// // 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// // 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// // 		}

// // 		return env.JWT_SECRET, nil
// // 	})

// // 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// // 		fmt.Println(claims["email"], claims["birthday"])
// // 	} else {
// // 		fmt.Println(err)
// // 	}
// // }
