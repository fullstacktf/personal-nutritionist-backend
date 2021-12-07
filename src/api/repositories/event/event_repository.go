package repositories

import (
	"errors"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepository struct {
	db *mongo.Database
}

func NewEventRepository(db *mongo.Database) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) GetEvents(c *gin.Context) ([]models.Event, error) {
	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("events")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []models.Event
	err = cursor.All(ctx, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *EventRepository) GetEventByID(c *gin.Context, id primitive.ObjectID) (*models.Event, error) {
	var event models.Event

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("events")
	result := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	if result == nil {
		return nil, errors.New("failed to find an user")
	}

	err := result.Decode(&event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
