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

func InitConnection() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(getConnectionURI()))
	if err != nil {
		log.Fatalln("Failed to create client 💣:", err)
	}

	ctx, cancel := GetContext(client)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln("Failed to connect to database 💣:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln("Failed to ping to database 💣:", err)
	}

	log.Println("Connected to MongoDB! 😊🍃")
	return client.Database(env.MONGO_DATABASE)
}

func GetContext(client *mongo.Client) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	return ctx, cancel
}

func DropConnection(db *mongo.Database, ctx context.Context, cancel context.CancelFunc) {
	cancel()
	db.Drop(ctx)
}
