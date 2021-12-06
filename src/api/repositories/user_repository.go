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

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUsers(c *gin.Context) ([]models.User, error) {
	client, ctx, cancel, collection := database.GetCollection("users")
	defer database.DropConnection(client, ctx, cancel)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(c *gin.Context, id primitive.ObjectID) (models.User, error) {
	var user models.User

	client, ctx, cancel, collection := database.GetCollection("users")
	defer database.DropConnection(client, ctx, cancel)
	result := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	if result == nil {
		return user, errors.New("failed to find an user")
	}

	err := result.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) PostUser(c *gin.Context, user *models.User) (primitive.ObjectID, error) {
	user.ObjectID = primitive.NewObjectID()

	client, ctx, cancel, collection := database.GetCollection("users")
	defer database.DropConnection(client, ctx, cancel)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	objectID := result.InsertedID.(primitive.ObjectID)

	return objectID, nil
}

func (r *UserRepository) PutUser(c *gin.Context, id primitive.ObjectID, newUser models.User) (models.User, error) {
	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.M{"$set": newUser}
	var user models.User

	client, ctx, cancel, collection := database.GetCollection("users")
	defer database.DropConnection(client, ctx, cancel)
	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (r *UserRepository) DeleteUser(c *gin.Context, id primitive.ObjectID) (models.User, error) {
	var user models.User

	client, ctx, cancel, collection := database.GetCollection("users")
	defer database.DropConnection(client, ctx, cancel)
	result := collection.FindOneAndDelete(ctx, bson.D{{Key: "_id", Value: id}})
	if result == nil {
		return user, errors.New("failed to delete an user")
	}

	err := result.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}