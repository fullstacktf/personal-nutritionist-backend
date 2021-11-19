package routes

import (
	"bytes"
	"github.com/fullstacktf/personal-nutritionist-backend/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	t.Run("should return users", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		context := gin.New()
		context.GET("/api/users/", handlers.GetUsers())

		req, err := http.NewRequest(http.MethodGet, "/api/users/", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		context.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
