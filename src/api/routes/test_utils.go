package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	eventRepositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/event"
	recipeRepositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/recipe"
	userRepositories "github.com/fullstacktf/personal-nutritionist-backend/api/repositories/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type Error struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

var (
	context              *gin.Engine
	userRepositoryMock   *userRepositories.UserRepositoryMock
	eventRepositoryMock  *eventRepositories.EventRepositoryMock
	recipeRepositoryMock *recipeRepositories.RecipeRepositoryMock
)

func setUp() {
	gin.SetMode(gin.TestMode)
	userRepositoryMock = new(userRepositories.UserRepositoryMock)
	eventRepositoryMock = new(eventRepositories.EventRepositoryMock)
	recipeRepositoryMock = new(recipeRepositories.RecipeRepositoryMock)
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
