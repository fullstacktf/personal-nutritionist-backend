package models

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	User        *User  `json:"user"`
	TokenString string `json:"token"`
}
