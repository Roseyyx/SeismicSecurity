package utilities

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectDB function
func ConnectDB() *mongo.Client {
	// Connect to MongoDB
	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(GetEnvVariable("DB_URL")))

	// If there is an error connecting to MongoDB
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	defer func() {
		// Disconnect from MongoDB
		if err := Client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Return the MongoDB client
	return Client
}
