package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ObjectID     primitive.ObjectID `json:"_id" bson:"_id"`
	Title        string             `json:"title" binding:"required"`
	Description  string             `json:"description" bson:"description"`
	Status       string             `json:"status" bson:"status"`
	Participants []BasicUser        `json:"participants" bson:"participants"`
	StartingDate string             `json:"startingDate" bson:"startingDate" binding:"required"`
	EndingDate   string             `json:"endingDate" bson:"endingDate" binding:"required"`
}
