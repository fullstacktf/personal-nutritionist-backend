package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectTimeout = 5

func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file 💣")
	}

	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	database := os.Getenv("MONGO_DATABASE")

	connectionURI := "mongodb://" + username + ":" + password + "@localhost/27017/" + database + "?authSource=admin"

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client 💣: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to database 💣: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping to database 💣: %v", err)
	}

	fmt.Println("Connected to MongoDB! 😊🍃")
	return client, ctx, cancel
}
