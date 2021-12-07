package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var recipesMock = []models.Recipe{
	{ObjectID: primitive.NewObjectID(), Name: "Porridge de nueces y arandanos", TypeDiet: "Vegana", TypeMeal: "Desayuno", Alergens: []string{"Frutos secos"}, Date: "24-12-2021", Ingredients: []string{"Nueces", "Arándanos"}},
	{ObjectID: primitive.NewObjectID(), Name: "Papas rellenas", TypeMeal: "Almuerzo", Alergens: []string{"Carne de cerdo", "Carne de vaca"}, Date: "12-12-2021", Ingredients: []string{"Papas", "Carne", "Mojo"}},
}

var recipeErrorMock = Error{Message: "error de receta", Status: "💣"}

func TestGetRecipes(t *testing.T) {
	t.Run("should return status OK and recipes", func(t *testing.T) {
		setUp()
		recipeRepositoryMock.On("GetRecipes", mock.AnythingOfType("*gin.Context")).Return(recipesMock, nil)
		context.GET("/api/users/:id/weekmeal/", handlers.GetRecipes(recipeRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/:id/weekmeal/", "")
		formerBody, err := json.MarshalIndent(recipesMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		recipeRepositoryMock.On("GetRecipes", mock.AnythingOfType("*gin.Context")).Return([]models.Recipe{}, errors.New("error de receta"))
		context.GET("/api/users/:id/weekmeal/", handlers.GetRecipes(recipeRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/:id/weekmeal/", "")
		formerBody, err := json.MarshalIndent(recipeErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestGetRecipeByID(t *testing.T) {
	t.Run("should return status OK and recipe", func(t *testing.T) {
		setUp()
		recipeRepositoryMock.On("GetRecipeByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&recipesMock[0], nil)
		context.GET("/api/users/:id/weekmeal/recipe/:idRecipe/", handlers.GetRecipeByID(recipeRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/:id/weekmeal/recipe/:idRecipe/", "")
		formerBody, err := json.MarshalIndent(recipesMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		recipeRepositoryMock.On("GetRecipeByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&models.Recipe{}, errors.New("error de receta"))
		context.GET("/api/users/:id/weekmeal/recipe/:idRecipe/", handlers.GetRecipeByID(recipeRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/:id/weekmeal/recipe/:idRecipe/", "")
		formerBody, err := json.MarshalIndent(recipeErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestCreateRecipe(t *testing.T) {
	t.Run("should return status OK and recipe", func(t *testing.T) {
		setUp()
		recipeRepositoryMock.On("CreateRecipe", mock.AnythingOfType("*gin.Context"), &recipesMock[0]).Return(recipesMock[0].ObjectID, nil)
		context.POST("/api/users/:id/weekmeal/recipe/", handlers.CreateRecipe(recipeRepositoryMock))

		reqBody, err := json.Marshal(recipesMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPost, "/api/users/:id/weekmeal/recipe/", string(reqBody))

		expect := "\"" + recipesMock[0].ObjectID.Hex() + "\""

		assert.Equal(t, http.StatusCreated, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, expect, rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		recipeRepositoryMock.On("CreateRecipe", mock.AnythingOfType("*gin.Context"), &recipesMock[0]).Return(primitive.NilObjectID, errors.New("error de receta"))
		context.POST("/api/users/:id/weekmeal/recipe/", handlers.CreateRecipe(recipeRepositoryMock))

		reqBody, err := json.Marshal(recipesMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPost, "/api/users/:id/weekmeal/recipe/", string(reqBody))

		formerBody, err := json.MarshalIndent(recipeErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestUpdateRecipe(t *testing.T) {
	t.Run("should return status OK and recipe", func(t *testing.T) {
		setUp()
		recipeRepositoryMock.On("UpdateRecipe", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID, &recipesMock[0]).Return(&recipesMock[0], nil)
		context.PUT("/api/users/:id/weekmeal/recipe/:idRecipe/", handlers.UpdateRecipe(recipeRepositoryMock))

		reqBody, err := json.Marshal(recipesMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPut, "/api/users/:id/weekmeal/recipe/:idRecipe/", string(reqBody))

		formerBody, err := json.MarshalIndent(recipesMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		recipeRepositoryMock.On("UpdateRecipe", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID, &recipesMock[0]).Return(&models.Recipe{}, errors.New("error de receta"))
		context.PUT("/api/users/:id/weekmeal/recipe/:idRecipe/", handlers.UpdateRecipe(recipeRepositoryMock))

		reqBody, err := json.Marshal(recipesMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPut, "/api/users/:id/weekmeal/recipe/:idRecipe/", string(reqBody))

		formerBody, err := json.MarshalIndent(recipeErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}
