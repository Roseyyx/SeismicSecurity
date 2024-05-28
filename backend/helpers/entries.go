package helpers

import (
	"encoding/json"
	"main/backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEntry(username, password, website, notes string) {
	// Create a new entry in the database file
	AddToFile(models.Entry{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: password,
		Website:  website,
		Notes:    notes,
	})
}

// return model.Entry or error
func GetEntries(password, filename string) []models.Entry {
	bytes := ReadFile(password)
	entries := []models.Entry{}
	_ = json.Unmarshal(bytes, &entries)
	return entries
}
