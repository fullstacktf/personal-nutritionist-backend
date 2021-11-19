package models

type Event struct {
	ID           string      `json:"id" binding:"required"`
	Title        string      `json:"title" binding:"required"`
	Description  string      `json:"description"`
	Status       string      `json:"status"`
	Participants []BasicUser `json:"participants"`
	StartingDate string      `json:"startingDate" binding:"required"`
	EndingDate   string      `json:"endingDate" binding:"required"`
}
