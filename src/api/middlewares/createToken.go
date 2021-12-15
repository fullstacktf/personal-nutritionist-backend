package middlewares

// import (
// 	"fmt"
// 	"time"
// )

// type MyCustomClaims struct {
// 	Email    string `json:"email"`
// 	Birthday int64  `json:"birthday"`
// 	jwt.StandardClaims
// }

// func createToken() {

// 	claims := MyCustomClaims{
// 		"hello@friendsofgo.tech",
// 		time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC).Unix(),
// 		jwt.StandardClaims{
// 			ExpiresAt: 15000,
// 			Issuer:    "Friends of Go",
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	mySecret := "my-secret"
// 	signedToken, err := token.SignedString(mySecret)

// }

// func parseToken(receivedToken) {
// 	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return mySecret, nil
// 	})
// }
