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

var eventsMock = []models.Event{
	{ObjectID: primitive.NewObjectID(), Title: "Reunión con Dolores", Owner: primitive.NewObjectID(), Description: "Reunión para controlar tiroides", Participants: []models.BasicUser{}, Status: "Confirmado", StartingDate: "12/12/2021 13:00", EndingDate: "12/12/2021 13:55"},
	{ObjectID: primitive.NewObjectID(), Title: "Reunión con Casiano", Owner: primitive.NewObjectID(), Description: "Reunión para controlar adicción al azúcar", Participants: []models.BasicUser{}, Status: "Pendiente", StartingDate: "12/12/2021 13:00", EndingDate: "12/12/2021 13:55"},
	{ObjectID: primitive.NewObjectID(), Title: "Reunión con Coromoto", Owner: primitive.NewObjectID(), Description: "Terapia trastorno alimenticio", Participants: []models.BasicUser{}, Status: "Confirmado", StartingDate: "12/12/2021 13:00", EndingDate: "12/12/2021 13:55"},
}

var eventErrorMock = Error{Message: "error de evento", Status: "💣"}

func TestGetEvents(t *testing.T) {
	t.Run("should return status OK and events", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("GetEvents", mock.AnythingOfType("*gin.Context")).Return(eventsMock, nil)
		context.GET("/calendar/users/:id", handlers.GetEvents(eventRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/calendar/users/:id", "")
		formerBody, err := json.MarshalIndent(eventsMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("GetEvents", mock.AnythingOfType("*gin.Context")).Return([]models.Event{}, errors.New("error de evento"))
		context.GET("/calendar/users/:id", handlers.GetEvents(eventRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/calendar/users/:id", "")
		formerBody, err := json.MarshalIndent(eventErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestGetEventByID(t *testing.T) {
	t.Run("should return status OK and event", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("GetEventByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&eventsMock[0], nil)
		context.GET("/calendar/event/:idEvent", handlers.GetEventByID(eventRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/calendar/event/:idEvent", "")
		formerBody, err := json.MarshalIndent(eventsMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("GetEventByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&models.Event{}, errors.New("error de evento"))
		context.GET("/calendar/event/:idEvent", handlers.GetEventByID(eventRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/calendar/event/:idEvent", "")
		formerBody, err := json.MarshalIndent(eventErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestCreateEvent(t *testing.T) {
	t.Run("should return status OK and event", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("CreateEvent", mock.AnythingOfType("*gin.Context"), &eventsMock[0]).Return(&eventsMock[0], nil)
		context.POST("/calendar/event", handlers.CreateEvent(eventRepositoryMock))

		reqBody, err := json.Marshal(eventsMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPost, "/calendar/event", string(reqBody))

		formerBody, err := json.MarshalIndent(eventsMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusCreated, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("CreateEvent", mock.AnythingOfType("*gin.Context"), &eventsMock[0]).Return(&models.Event{}, errors.New("error de evento"))
		context.POST("/calendar/event", handlers.CreateEvent(eventRepositoryMock))

		reqBody, err := json.Marshal(eventsMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPost, "/calendar/event", string(reqBody))

		formerBody, err := json.MarshalIndent(eventErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestUpdateEvent(t *testing.T) {
	t.Run("should return status OK and event", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("UpdateEvent", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID, &eventsMock[0]).Return(&eventsMock[0], nil)
		context.PUT("/calendar/event/:idEvent", handlers.UpdateEvent(eventRepositoryMock))

		reqBody, err := json.Marshal(eventsMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPut, "/calendar/event/:idEvent", string(reqBody))

		formerBody, err := json.MarshalIndent(eventsMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("UpdateEvent", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID, &eventsMock[0]).Return(&models.Event{}, errors.New("error de evento"))
		context.PUT("/calendar/event/:idEvent", handlers.UpdateEvent(eventRepositoryMock))

		reqBody, err := json.Marshal(eventsMock[0])
		require.NoError(t, err)
		res, rec := executeRequest(t, http.MethodPut, "/calendar/event/:idEvent", string(reqBody))

		formerBody, err := json.MarshalIndent(eventErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}

func TestDeleteEvent(t *testing.T) {
	t.Run("should return status OK and event", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("DeleteEvent", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&eventsMock[0], nil)
		context.DELETE("/calendar/event/:idEvent", handlers.DeleteEvent(eventRepositoryMock))

		res, rec := executeRequest(t, http.MethodDelete, "/calendar/event/:idEvent", "")
		formerBody, err := json.MarshalIndent(eventsMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("DeleteEvent", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&models.Event{}, errors.New("error de evento"))
		context.DELETE("/calendar/event/:idEvent", handlers.DeleteEvent(eventRepositoryMock))

		res, rec := executeRequest(t, http.MethodDelete, "/calendar/event/:idEvent", "")
		formerBody, err := json.MarshalIndent(eventErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}
