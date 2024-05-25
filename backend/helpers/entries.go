package helpers

import (
	"log"
	"main/backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEntry(username, password, website, notes string) {
	// Encrypt the password we got from the file
	//encryptedPassword := utilities.Encrypt([]byte(GetFromMemory().Password), utilities.CreateHash(password))

	// Create a new entry in the database file
	err := AddToFile(models.Entry{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: password,
		Website:  website,
		Notes:    notes,
	}, GetFromMemory().Filename)

	if err != nil {
		log.Println("Error adding entry to file: ", err)
	}

}

// return model.Entry or error
func GetEntries(password, filename string) []models.Entry {
	// Get the entries in bytes from the file
	fileBytes := OpenFile(filename, password)

	// Get the entries from the bytes
	entries, err := GetEntriesFromBytes(fileBytes)
	if err != nil {
		log.Println("Error getting entries from bytes: ", err)
	}

	return entries
}
