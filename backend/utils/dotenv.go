package utilities

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Environment variables function
func GetEnvVariable(key string) string {
	// Load .env file
	err := godotenv.Load(".env")

	// If there is an error loading the .env file
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Return the environment variable
	return os.Getenv(key)
}
