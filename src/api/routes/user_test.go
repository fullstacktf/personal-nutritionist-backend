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

var userMock = models.User{ObjectID: primitive.NewObjectID(), Name: "Sergio Peinado", Email: "sergiopeinado@gmail.com", Role: "Nutritionist", Username: "Sergito", Dni: "41257854L", Phone: 612732894, Likes: 157770, IsVerified: true, Password: "1234"}

type ErrorMock struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

var errorMock = ErrorMock{Message: "error de prueba", Status: "💣"}

func TestGetUsers(t *testing.T) {
	t.Run("should return status OK and users", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		userRepositoryMock := new(repositories.UserRepositoryMock)
		c := gin.New()

		userRepositoryMock.On("GetUsers", mock.AnythingOfType("*gin.Context")).Return(usersMock, nil)
		c.GET("/api/users/", handlers.GetUsers(userRepositoryMock))

		req, err := http.NewRequest(http.MethodGet, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		c.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		formerBody, err := json.MarshalIndent(usersMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		userRepositoryMock := new(repositories.UserRepositoryMock)
		c := gin.New()

		userRepositoryMock.On("GetUsers", mock.AnythingOfType("*gin.Context")).Return([]models.User{}, errors.New("error de prueba"))
		c.GET("/api/users/", handlers.GetUsers(userRepositoryMock))

		req, err := http.NewRequest(http.MethodGet, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		c.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		formerBody, err := json.MarshalIndent(errorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		userRepositoryMock := new(repositories.UserRepositoryMock)
		c := gin.New()

		userRepositoryMock.On("GetUserByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(usersMock[0], nil)
		c.GET("/api/users/", handlers.GetUserByID(userRepositoryMock))

		req, err := http.NewRequest(http.MethodGet, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		c.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		formerBody, err := json.MarshalIndent(usersMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		userRepositoryMock := new(repositories.UserRepositoryMock)
		c := gin.New()

		userRepositoryMock.On("GetUserByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(models.User{}, errors.New("error de prueba"))
		c.GET("/api/users/", handlers.GetUserByID(userRepositoryMock))

		req, err := http.NewRequest(http.MethodGet, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		c.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		formerBody, err := json.MarshalIndent(errorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestPostUser(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		userRepositoryMock := new(repositories.UserRepositoryMock)
		c := gin.New()

		userRepositoryMock.On("PostUser", mock.AnythingOfType("*gin.Context"), userMock).Return(usersMock[0].ObjectID, nil)
		c.POST("/api/users/", handlers.PostUser(userRepositoryMock))

		reqBody, _ := json.Marshal(userMock)
		req, err := http.NewRequest(http.MethodPost, "/api/users/", bytes.NewBuffer(reqBody))

		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		c.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, usersMock[0].ObjectID.Hex(), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		userRepositoryMock := new(repositories.UserRepositoryMock)
		c := gin.New()

		userRepositoryMock.On("GetUserByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(models.User{}, errors.New("error de prueba"))
		c.GET("/api/users/", handlers.GetUserByID(userRepositoryMock))

		req, err := http.NewRequest(http.MethodGet, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		c.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		formerBody, err := json.MarshalIndent(errorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestPutUser(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		userRepositoryMock := new(repositories.UserRepositoryMock)
		c := gin.New()

		userRepositoryMock.On("PutUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID, usersMock[0]).Return(usersMock[0], nil)
		c.PUT("/api/users/", handlers.PutUser(userRepositoryMock))

		req, err := http.NewRequest(http.MethodPut, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		c.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		formerBody, err := json.MarshalIndent(usersMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		userRepositoryMock := new(repositories.UserRepositoryMock)
		c := gin.New()

		userRepositoryMock.On("GetUserByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(models.User{}, errors.New("error de prueba"))
		c.GET("/api/users/", handlers.GetUserByID(userRepositoryMock))

		req, err := http.NewRequest(http.MethodGet, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		c.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		formerBody, err := json.MarshalIndent(errorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		userRepositoryMock := new(repositories.UserRepositoryMock)
		c := gin.New()

		userRepositoryMock.On("DeleteUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(usersMock[0], nil)
		c.DELETE("/api/users/", handlers.DeleteUser(userRepositoryMock))

		req, err := http.NewRequest(http.MethodDelete, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		c.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		formerBody, err := json.MarshalIndent(usersMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		userRepositoryMock := new(repositories.UserRepositoryMock)
		c := gin.New()

		userRepositoryMock.On("DeleteUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(models.User{}, errors.New("error de prueba"))
		c.DELETE("/api/users/", handlers.DeleteUser(userRepositoryMock))

		req, err := http.NewRequest(http.MethodDelete, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		c.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		formerBody, err := json.MarshalIndent(errorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}
