package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ObjectID     primitive.ObjectID `json:"_id" bson:"_id"`
	Title        string             `json:"title" bson:"title" binding:"required"`
	Description  string             `json:"description" bson:"description"`
	Owner        primitive.ObjectID `json:"owner" bson:"owner" binding:"required"`
	Status       string             `json:"status" bson:"status"`
	Participants []BasicUser        `json:"participants" bson:"participants" binding:"required"`
	StartingDate string             `json:"startingDate" bson:"startingDate" binding:"required"`
	EndingDate   string             `json:"endingDate" bson:"endingDate" binding:"required"`
}
