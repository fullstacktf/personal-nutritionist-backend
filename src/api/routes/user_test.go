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

var usersMock = []models.User{
	{ObjectID: primitive.NewObjectID(), Name: "Sergio Peinado", Email: "sergiopeinado@gmail.com", Role: "Nutritionist", Username: "Sergito", Dni: "41257854L", Phone: 612732894, Likes: 157770, IsVerified: true, Password: "1234"},
	{ObjectID: primitive.NewObjectID(), Name: "Godhito", Email: "damecomidah@gmail.com", Dni: "87654321P", Username: "Adanito", TypeDiet: "Hypercaloric", Weight: 120, Height: 160, Role: "Client", Password: "1234"},
	{ObjectID: primitive.NewObjectID(), Name: "Sarah Vaughan", Dni: "12345678P", TypeDiet: "vegetarian", Weight: 60, Height: 173, Role: "Client", Password: "1234"},
}
var credentialMock = models.Auth{Email: "sergiopeinado@gmail.com", Password: "1234"}
var tokenMock = models.Token{Email: "sergiopeinado@gmail.com", Role: "Nutricionista", TokenString: "token de prueba"}
var userErrorMock = Error{Message: "error de usuario", Status: "💣"}

func TestSignUp(t *testing.T) {
	t.Run("should return status OK and message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("SignUp", mock.AnythingOfType("*gin.Context"), &usersMock[0]).Return(&tokenMock, nil)
		context.POST("/auth/signup", handlers.SignUp(userRepositoryMock))

		reqBody, err := json.Marshal(usersMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPost, "/auth/signup", string(reqBody))

		formerBody, err := json.MarshalIndent(tokenMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusCreated, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("SignUp", mock.AnythingOfType("*gin.Context"), &usersMock[0]).Return(&models.Token{}, errors.New("error de usuario"))
		context.POST("/auth/signup", handlers.SignUp(userRepositoryMock))

		reqBody, err := json.Marshal(usersMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPost, "/auth/signup", string(reqBody))

		formerBody, err := json.MarshalIndent(userErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestLogIn(t *testing.T) {
	t.Run("should return status OK and message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("LogIn", mock.AnythingOfType("*gin.Context"), &credentialMock).Return(&tokenMock, nil)
		context.POST("/auth/login", handlers.LogIn(userRepositoryMock))

		reqBody, err := json.Marshal(credentialMock)
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPost, "/auth/login", string(reqBody))

		formerBody, err := json.MarshalIndent(tokenMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("LogIn", mock.AnythingOfType("*gin.Context"), &credentialMock).Return(&models.Token{}, errors.New("error de usuario"))
		context.POST("/auth/login", handlers.LogIn(userRepositoryMock))

		reqBody, err := json.Marshal(credentialMock)
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPost, "/auth/login", string(reqBody))

		formerBody, err := json.MarshalIndent(userErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestGetUsers(t *testing.T) {
	t.Run("should return status OK and users", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("GetUsers", mock.AnythingOfType("*gin.Context")).Return(usersMock, nil)
		context.GET("/users", handlers.GetUsers(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/users", "")
		formerBody, err := json.MarshalIndent(usersMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("GetUsers", mock.AnythingOfType("*gin.Context")).Return([]models.User{}, errors.New("error de usuario"))
		context.GET("/users", handlers.GetUsers(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/users", "")
		formerBody, err := json.MarshalIndent(userErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("GetUserByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&usersMock[0], nil)
		context.GET("/users/:id", handlers.GetUserByID(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/users/:id", "")
		formerBody, err := json.MarshalIndent(usersMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("GetUserByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&models.User{}, errors.New("error de usuario"))
		context.GET("/users/:id", handlers.GetUserByID(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/users/:id", "")
		formerBody, err := json.MarshalIndent(userErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestGetUsersByRole(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("GetUsersByRole", mock.AnythingOfType("*gin.Context"), ":role").Return(usersMock, nil)
		context.GET("/users/role/:role", handlers.GetUsersByRole(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/users/role/:role", "")
		formerBody, err := json.MarshalIndent(usersMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("GetUsersByRole", mock.AnythingOfType("*gin.Context"), ":role").Return([]models.User{}, errors.New("error de usuario"))
		context.GET("/users/role/:role", handlers.GetUsersByRole(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/users/role/:role", "")
		formerBody, err := json.MarshalIndent(userErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("UpdateUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID, &usersMock[0]).Return(&usersMock[0], nil)
		context.PUT("/users/:id", handlers.UpdateUser(userRepositoryMock))

		reqBody, err := json.Marshal(usersMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPut, "/users/:id", string(reqBody))

		formerBody, err := json.MarshalIndent(usersMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("UpdateUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID, &usersMock[0]).Return(&models.User{}, errors.New("error de usuario"))
		context.PUT("/users/:id", handlers.UpdateUser(userRepositoryMock))

		reqBody, err := json.Marshal(usersMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPut, "/users/:id", string(reqBody))

		formerBody, err := json.MarshalIndent(userErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("should return status OK and user", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("DeleteUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&usersMock[0], nil)
		context.DELETE("/users/:id", handlers.DeleteUser(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodDelete, "/users/:id", "")
		formerBody, err := json.MarshalIndent(usersMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		userRepositoryMock.On("DeleteUser", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&models.User{}, errors.New("error de usuario"))
		context.DELETE("/users/:id", handlers.DeleteUser(userRepositoryMock))

		res, rec := executeRequest(t, http.MethodDelete, "/users/:id", "")
		formerBody, err := json.MarshalIndent(userErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}
