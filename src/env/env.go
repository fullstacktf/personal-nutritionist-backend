package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MONGO_USERNAME string
var MONGO_PASSWORD string
var MONGO_DATABASE string
var MONGO_URL string
var MONGO_PORT string

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file 💣")
	}

	MONGO_USERNAME = os.Getenv("MONGO_USERNAME")
	MONGO_PASSWORD = os.Getenv("MONGO_PASSWORD")
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	MONGO_URL = os.Getenv("MONGO_URL")
	MONGO_PORT = os.Getenv("MONGO_PORT")

	log.Println("Environment variables loaded! 😊🪄")
}
