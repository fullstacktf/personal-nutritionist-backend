package repositories

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeRepository struct {
	db *mongo.Database
}

func NewRecipeRepository(db *mongo.Database) models.RecipeRepository {
	return &RecipeRepository{
		db: db,
	}
}

func (r *RecipeRepository) GetRecipes(c *gin.Context) ([]models.Recipe, error) {
	ctx, cancel := database.GetContext(r.db.Client())
	defer database.DropConnection(r.db, ctx, cancel)

	collection := r.db.Collection("recipes")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var recipes []models.Recipe
	err = cursor.All(ctx, &recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
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
