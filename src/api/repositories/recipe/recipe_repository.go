package repositories

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RecipeRepository struct {
	db *mongo.Database
}

func NewRecipeRepository(db *mongo.Database) models.RecipeRepository {
	return &RecipeRepository{
		db: db,
	}
}
func (r *RecipeRepository) CreateRecipe(c *gin.Context, recipe *models.Recipe) (primitive.ObjectID, error) {
	recipe.ObjectID = primitive.NewObjectID()

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("recipes")
	result, err := collection.InsertOne(ctx, recipe)
	if err != nil {
		return primitive.NilObjectID, err
	}
	objectID := result.InsertedID.(primitive.ObjectID)

	return objectID, nil
}

func (r *RecipeRepository) UpdateRecipe(c *gin.Context, id primitive.ObjectID, newRecipe *models.Recipe) (*models.Recipe, error) {
	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.M{"$set": newRecipe}
	var recipe models.Recipe

	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("recipes")
	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&recipe)
	if err != nil {
		return nil, err
	}

	return newRecipe, nil
}
