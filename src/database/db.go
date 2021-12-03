package database

import (
	"context"
	"log"
	"time"

	"github.com/fullstacktf/personal-nutritionist-backend/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectTimeout = 5

func getConnectionURI() string {
	env.LoadEnv()
	return "mongodb://" + env.MONGO_USERNAME + ":" + env.MONGO_PASSWORD + "@" + env.MONGO_URL + "/" + env.MONGO_PORT + "/" + env.MONGO_DATABASE + "?authSource=admin"
}

func InitConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(getConnectionURI()))
	if err != nil {
		log.Fatalln("Failed to create client 💣:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln("Failed to connect to database 💣:", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln("Failed to ping to database 💣:", err)
	}

	// log.Println("----------", client.Database("nutriguide"))
	log.Println("Connected to MongoDB! 😊🍃")
	return client, ctx, cancel
}

func GetCollection(collectionName string) (*mongo.Client, context.Context, context.CancelFunc, *mongo.Collection) {
	client, ctx, cancel := InitConnection()

	collection := client.Database("nutriguide").Collection(collectionName)
	return client, ctx, cancel, collection
}

func DropConnection(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	cancel()
	client.Disconnect(ctx)
}
