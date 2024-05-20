package utilities

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var HasDatabaseConnection bool

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

	// Check the connection
	err = Client.Ping(context.Background(), nil)

	// If there is an error checking the connection
	if err != nil {
		log.Fatalf("Error checking the connection: %v", err)
	} else {
		HasDatabaseConnection = true
	}

	log.Println("Connected to MongoDB")

	// Return the MongoDB client
	return Client
}
