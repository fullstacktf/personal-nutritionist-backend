package repositories

import (
	"database/sql"
	"github.com/fullstacktf/personal-nutritionist-backend/api/models"
	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: "1", Name: "Sergio Peinado", Email: "sergiopeinado@gmail.com", Role: "Nutritionist", Username: "Sergito", Dni: "41257854L", Phone: 612732894, Likes: 157770, IsVerified: true},
	{ID: "2", Name: "Godhito", Email: "damecomidah@gmail.com", Dni: "87654321P", Username: "Adanito", TypeDiet: "Hypercaloric", Weight: 120, Height: 160, Role: "Client"},
	{ID: "3", Name: "Sarah Vaughan", Dni: "12345678P", TypeDiet: "vegetarian", Weight: 60, Height: 173, Role: "Client"},
}

type UserMongoRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserMongoRepository {
	return &UserMongoRepository{
		db: db,
	}
}

func (r *UserMongoRepository) GetUsers(context *gin.Context) []models.User {
	return users
}
