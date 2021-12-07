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
	{ObjectID: primitive.NewObjectID(), Title: "Reunión con Dolores", Description: "Reunión para controlar tiroides", Status: "Confirmado", StartingDate: "12/12/2021 13:00", EndingDate: "12/12/2021 13:55"},
	{ObjectID: primitive.NewObjectID(), Title: "Reunión con Casiano", Description: "Reunión para controlar adicción al azúcar", Status: "Pendiente", StartingDate: "12/12/2021 13:00", EndingDate: "12/12/2021 13:55"},
	{ObjectID: primitive.NewObjectID(), Title: "Reunión con Coromoto", Description: "Terapia trastorno alimenticio", Status: "Confirmado", StartingDate: "12/12/2021 13:00", EndingDate: "12/12/2021 13:55"},
}

var eventErrorMock = Error{Message: "error de evento", Status: "💣"}

func TestGetEvents(t *testing.T) {
	t.Run("should return status OK and events", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("GetEvents", mock.AnythingOfType("*gin.Context")).Return(eventsMock, nil)
		context.GET("/api/users/:id/calendar/", handlers.GetEvents(eventRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/:id/calendar/", "")
		formerBody, err := json.MarshalIndent(eventsMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("GetEvents", mock.AnythingOfType("*gin.Context")).Return(eventsMock, errors.New("error de evento"))
		context.GET("/api/users/:id/calendar/", handlers.GetEvents(eventRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/:id/calendar/", "")
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
		context.GET("/api/users/:id/calendar/event/", handlers.GetEventByID(eventRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/:id/calendar/event/", "")
		formerBody, err := json.MarshalIndent(eventsMock[0], "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})

	t.Run("should return error status and error message", func(t *testing.T) {
		setUp()
		eventRepositoryMock.On("GetEventByID", mock.AnythingOfType("*gin.Context"), primitive.NilObjectID).Return(&eventsMock[0], errors.New("error de evento"))
		context.GET("/api/users/:id/calendar/event/", handlers.GetEventByID(eventRepositoryMock))

		res, rec := executeRequest(t, http.MethodGet, "/api/users/:id/calendar/event/", "")
		formerBody, err := json.MarshalIndent(eventErrorMock, "", "    ")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode, "they should be equal 💣")
		assert.Equal(t, string(formerBody), rec.Body.String(), "they should be equal 💣")
	})
}
