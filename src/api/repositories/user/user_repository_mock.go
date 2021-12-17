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

func (m *UserRepositoryMock) SignUp(c *gin.Context, user *models.User) (*models.Token, error) {
	args := m.Called(c, user)
	return args.Get(0).(*models.Token), args.Error(1)
}

func (m *UserRepositoryMock) LogIn(c *gin.Context, credential *models.Auth) (*models.Token, error) {
	args := m.Called(c, credential)
	return args.Get(0).(*models.Token), args.Error(1)
}

func (m *UserRepositoryMock) GetUsers(c *gin.Context) ([]models.User, error) {
	args := m.Called(c)
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUsersByRole(c *gin.Context, role string) ([]models.User, error) {
	args := m.Called(c, role)
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByID(c *gin.Context, id primitive.ObjectID) (*models.User, error) {
	args := m.Called(c, id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) UpdateUser(c *gin.Context, id primitive.ObjectID, newUser *models.User) (*models.User, error) {
	args := m.Called(c, id, newUser)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) DeleteUser(c *gin.Context, id primitive.ObjectID) (*models.User, error) {
	args := m.Called(c, id)
	return args.Get(0).(*models.User), args.Error(1)
}
