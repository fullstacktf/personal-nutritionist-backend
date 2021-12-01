package repositories

import (
	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetUsers(c *gin.Context) ([]models.User, error) {
	args := m.Called(c)
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByID(c *gin.Context, id primitive.ObjectID) (models.User, error) {
	var user models.User
	return user, nil
}

func (m *UserRepositoryMock) PostUser(c *gin.Context, user *models.User) (primitive.ObjectID, error) {
	return primitive.NewObjectID(), nil
}

func (m *UserRepositoryMock) PutUser(c *gin.Context, id primitive.ObjectID, newUser models.User) (models.User, error) {
	return newUser, nil
}

func (m *UserRepositoryMock) DeleteUser(c *gin.Context, id primitive.ObjectID) (models.User, error) {
	var user models.User
	return user, nil
}
