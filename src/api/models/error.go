package models

type Error struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
