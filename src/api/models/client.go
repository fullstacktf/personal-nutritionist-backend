package models

type Client struct {
	ID           string   `json:"id" binding:"required"`
	Name         string   `json:"name" binding:"required"`
	Dni          string   `json:"dni"`
	Phone        uint32   `json:"phone"`
	Type_diet    string   `json:"type_diet"`
	Weight       uint     `json:"weight"`
	Height       uint     `json:"height"`
	Birthdate    string   `json:"birthdate"`
	Intolerances []string `json:"intolerances"`
	Description  string   `json:"description"`
	Photo        string   `json:"photo"`
}
