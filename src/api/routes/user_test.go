package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	repositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var usersMock = []models.User{
	{ObjectID: primitive.NewObjectID(), Name: "Sergio Peinado", Email: "sergiopeinado@gmail.com", Role: "Nutritionist", Username: "Sergito", Dni: "41257854L", Phone: 612732894, Likes: 157770, IsVerified: true, Password: "1234"},
	{ObjectID: primitive.NewObjectID(), Name: "Godhito", Email: "damecomidah@gmail.com", Dni: "87654321P", Username: "Adanito", TypeDiet: "Hypercaloric", Weight: 120, Height: 160, Role: "Client", Password: "1234"},
	{ObjectID: primitive.NewObjectID(), Name: "Sarah Vaughan", Dni: "12345678P", TypeDiet: "vegetarian", Weight: 60, Height: 173, Role: "Client", Password: "1234"},
}

type ErrorMock struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

var errorMock = ErrorMock{Message: "error de prueba", Status: "💣"}

var (
	context            *gin.Engine
	userRepositoryMock *repositories.UserRepositoryMock
)

func setUp() {
	gin.SetMode(gin.TestMode)
	userRepositoryMock = new(repositories.UserRepositoryMock)
	context = gin.New()
}

func executeRequest(t *testing.T, method string, url string, request string) (*http.Response, *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(request)))
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	context.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	return res, rec
}

func TestGetUsers(t *testing.T) {
	t.Run("should return status OK and users", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("GetUsers", mock.AnythingOfType("*gin.Context")).Return(usersMock, nil)
		context.GET("/api/users/", handlers.GetUsers(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/", "")
		formerBody, err := json.MarshalIndent(usersMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("GetUsers", mock.AnythingOfType("*gin.Context")).Return([]models.User{}, errors.New("error de prueba"))
		context.GET("/api/users/", handlers.GetUsers(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/", "")
		formerBody, err := json.MarshalIndent(errorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("GetUserByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&usersMock[0], nil)
		context.GET("/api/users/", handlers.GetUserByID(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/", "")
		formerBody, err := json.MarshalIndent(usersMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("GetUserByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&models.User{}, errors.New("error de prueba"))
		context.GET("/api/users/", handlers.GetUserByID(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/", "")
		formerBody, err := json.MarshalIndent(errorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("CreateUser", mock.AnythingOfType("*gin.Context"), &usersMock[0]).Return(usersMock[0].ObjectID, nil)
		context.POST("/api/users/", handlers.CreateUser(userRepositoryMock))

		reqBody, err := json.Marshal(usersMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPost, "/api/users/", string(reqBody))

		expect := "\"" + usersMock[0].ObjectID.Hex() + "\""

		assert.Equal(t, http.StatusCreated, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, expect, rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("CreateUser", mock.AnythingOfType("*gin.Context"), &usersMock[0]).Return(primitive.NilObjectID, errors.New("error de prueba"))
		context.POST("/api/users/", handlers.CreateUser(userRepositoryMock))

		reqBody, err := json.Marshal(usersMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPost, "/api/users/", string(reqBody))

		formerBody, err := json.MarshalIndent(errorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("UpdateUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID, &usersMock[0]).Return(&usersMock[0], nil)
		context.PUT("/api/users/", handlers.UpdateUser(userRepositoryMock))

		reqBody, err := json.Marshal(usersMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPut, "/api/users/", string(reqBody))

		formerBody, err := json.MarshalIndent(usersMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("UpdateUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID, &usersMock[0]).Return(&usersMock[0], errors.New("error de prueba"))
		context.PUT("/api/users/", handlers.UpdateUser(userRepositoryMock))

		reqBody, err := json.Marshal(usersMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPut, "/api/users/", string(reqBody))

		formerBody, err := json.MarshalIndent(errorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("DeleteUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&usersMock[0], nil)
		context.DELETE("/api/users/", handlers.DeleteUser(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodDelete, "/api/users/", "")
		formerBody, err := json.MarshalIndent(usersMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("DeleteUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&models.User{}, errors.New("error de prueba"))
		context.DELETE("/api/users/", handlers.DeleteUser(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodDelete, "/api/users/", "")
		formerBody, err := json.MarshalIndent(errorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}
