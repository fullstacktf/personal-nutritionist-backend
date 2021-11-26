package routes

import (
	"bytes"
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/fullstacktf/personal-nutritionist-backend/api/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

var repositoryMock *repositories.UserRepositoryMock

var usersMock = []models.User{
	{ID: "1", Name: "Sergio Peinado", Email: "sergiopeinado@gmail.com", Role: "Nutritionist", Username: "Sergito", Dni: "41257854L", Phone: 612732894, Likes: 157770, IsVerified: true},
	{ID: "2", Name: "Godhito", Email: "damecomidah@gmail.com", Dni: "87654321P", Username: "Adanito", TypeDiet: "Hypercaloric", Weight: 120, Height: 160, Role: "Client"},
	{ID: "3", Name: "Sarah Vaughan", Dni: "12345678P", TypeDiet: "vegetarian", Weight: 60, Height: 173, Role: "Client"},
}

func TestGetUsers(t *testing.T) {
	t.Run("should return users", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		context := gin.New()

		repositoryMock.On("GetUsers", context).Return(usersMock, nil)
		context.GET("/api/users/", handlers.GetUsers(repositoryMock))

		req, err := http.NewRequest(http.MethodGet, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		context.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
