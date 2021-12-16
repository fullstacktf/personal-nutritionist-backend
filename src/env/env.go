package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MONGO_USERNAME string
	MONGO_PASSWORD string
	MONGO_DATABASE string
	MONGO_URL      string
	MONGO_PORT     string
	JWT_SECRET     string
)

func LoadEnv() {
	environment := os.Getenv("ENV")
	if environment == "local" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file 💣")
		}
	}

	MONGO_USERNAME = os.Getenv("MONGO_USERNAME")
	MONGO_PASSWORD = os.Getenv("MONGO_PASSWORD")
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	MONGO_URL = os.Getenv("MONGO_URL")
	MONGO_PORT = os.Getenv("MONGO_PORT")
	JWT_SECRET = os.Getenv("JWT_SECRET")

	log.Println("Environment variables loaded! 😊🪄")
}
