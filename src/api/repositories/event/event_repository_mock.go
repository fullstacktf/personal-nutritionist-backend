package repositories

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventRepositoryMock struct {
	mock.Mock
}

func (m *EventRepositoryMock) GetEvents(c *gin.Context) ([]models.Event, error) {
	args := m.Called(c)
	return args.Get(0).([]models.Event), args.Error(1)
}

func (m *EventRepositoryMock) CreateEvent(c *gin.Context, event *models.Event) (primitive.ObjectID, error) {
	args := m.Called(c, event)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}
