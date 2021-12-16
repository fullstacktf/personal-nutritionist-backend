package repositories

import (
	"errors"

	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/fullstacktf/personal-nutritionist-backend/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) models.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) SignUp(c *gin.Context, user *models.User) (*string, error) {
	newUser, err := r.GetUserByUsernameAndPassword(user.Email, user.Password)
	if err == nil {
		return nil, errors.New("user" + newUser.Email + "already exists")
	}

	user.ObjectID = primitive.NewObjectID()

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("users")
	if _, err := collection.InsertOne(ctx, user); err != nil {
		return nil, err
	}

	token, err := services.GenerateJWT(user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (r *UserRepository) LogIn(c *gin.Context, credential *models.Auth) (*string, error) {
	user, err := r.GetUserByUsernameAndPassword(credential.Email, credential.Password)
	if err != nil {
		return nil, err
	}

	token, err := services.GenerateJWT(user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (r *UserRepository) GetUsers(c *gin.Context) ([]models.User, error) {
	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("users")
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

func (r *UserRepository) GetUserByID(c *gin.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("users")
	result := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	if result == nil {
		return nil, errors.New("failed to find an user")
	}

	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUsersByRole(c *gin.Context, role string) ([]models.User, error) {
	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("users")
	cursor, err := collection.Find(ctx, bson.D{{Key: "role", Value: role}})
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

func (r *UserRepository) GetUserByUsernameAndPassword(email, password string) (*models.User, error) {
	var user models.User

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("users")
	result := collection.FindOne(ctx, bson.D{{Key: "email", Value: email}, {Key: "password", Value: password}})
	if result == nil {
		return nil, errors.New("failed to find an user")
	}

	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(c *gin.Context, id primitive.ObjectID, newUser *models.User) (*models.User, error) {
	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.M{"$set": newUser}
	var user models.User

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("users")
	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (r *UserRepository) DeleteUser(c *gin.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("users")
	result := collection.FindOneAndDelete(ctx, bson.D{{Key: "_id", Value: id}})
	if result == nil {
		return nil, errors.New("failed to delete an user")
	}

	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
