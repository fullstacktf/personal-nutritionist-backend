package repositories

import (
	"errors"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *EventRepository) CreateEvent(c *gin.Context, event *models.Event) (*models.Event, error) {
	event.ObjectID = primitive.NewObjectID()

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("events")
	if _, err := collection.InsertOne(ctx, event); err != nil {
		return nil, err
	}

	return event, nil
}

func (r *EventRepository) UpdateEvent(c *gin.Context, id primitive.ObjectID, newEvent *models.Event) (*models.Event, error) {
	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.M{"$set": newEvent}
	var event models.Event

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("events")
	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&event)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

func (r *EventRepository) DeleteEvent(c *gin.Context, id primitive.ObjectID) (*models.Event, error) {
	var event models.Event

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("events")
	result := collection.FindOneAndDelete(ctx, bson.D{{Key: "_id", Value: id}})
	if result == nil {
		return nil, errors.New("failed to delete an element")
	}

	err := result.Decode(&event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
