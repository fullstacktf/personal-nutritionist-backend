package models

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Email       string `json:"email"`
	Role        string `json:"role"`
	TokenString string `json:"token"`
}
