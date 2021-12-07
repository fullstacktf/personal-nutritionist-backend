package repositories

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeRepositoryMock struct {
	mock.Mock
}

func (m *RecipeRepositoryMock) GetRecipes(c *gin.Context) ([]models.Recipe, error) {
	args := m.Called(c)
	return args.Get(0).([]models.Recipe), args.Error(1)
}

func (m *RecipeRepositoryMock) GetRecipeByID(c *gin.Context, id primitive.ObjectID) (*models.Recipe, error) {
	args := m.Called(c, id)
	return args.Get(0).(*models.Recipe), args.Error(1)
}

func (m *RecipeRepositoryMock) CreateRecipe(c *gin.Context, recipe *models.Recipe) (primitive.ObjectID, error) {
	args := m.Called(c, recipe)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}
