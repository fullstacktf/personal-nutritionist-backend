package services

import (
	"errors"
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

func GenerateJWT(user *models.User) (*models.Token, error) {
	key := []byte(env.JWT_SECRET)
	claims := AuthClaims{
		true,
		user.Email,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(time.Minute*30).Unix(), 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}
	tokenStruct := models.Token{User: user, TokenString: signedToken}

	return &tokenStruct, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("invalid token")
		}
		return []byte(env.JWT_SECRET), nil
	})
}
