package helpers

import (
	"log"
	"main/backend/models"
	utilities "main/backend/utils"

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
func GetEntries(password string) []models.Entry {
	bytes := ReadFile(password)

	split := utilities.Split(string(bytes), "}")
	log.Println(bytes)

	entries := make([]models.Entry, 0)

	for _, entry := range split {
		if entry == "" {
			continue
		}

		entries = append(entries, models.Entry{
			ID:       primitive.NewObjectID(),
			Username: entry,
		})
	}

	return entries
}
