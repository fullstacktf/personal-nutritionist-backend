package repositories

import (
	"errors"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
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

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User

	client, ctx, cancel := database.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	cursor, err := client.Database("nutriguide").Collection("users").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(id primitive.ObjectID) (models.User, error) {
	var user models.User

	client, ctx, cancel := database.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	result := client.Database("nutriguide").Collection("users").FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	if result == nil {
		return user, errors.New("failed to find an user")
	}
	err := result.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) PostUser(user *models.User) (primitive.ObjectID, error) {
	client, ctx, cancel := database.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	user.ObjectID = primitive.NewObjectID()

	result, err := client.Database("nutriguide").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	objectId := result.InsertedID.(primitive.ObjectID)

	return objectId, nil
}

func (r *UserRepository) PutUser(id primitive.ObjectID, newUser models.User) (models.User, error) {
	var user models.User

	client, ctx, cancel := database.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.M{"$set": newUser}

	err := client.Database("nutriguide").Collection("users").FindOneAndUpdate(ctx, filter, update, opts).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) DeleteUser(id primitive.ObjectID) (models.User, error) {
	var user models.User

	client, ctx, cancel := database.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	result := client.Database("nutriguide").Collection("users").FindOneAndDelete(ctx, bson.D{{Key: "_id", Value: id}})
	if result == nil {
		return user, errors.New("failed to delete an user")
	}
	err := result.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
