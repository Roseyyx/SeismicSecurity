package helpers

import (
	"context"
	"log"
	"main/backend/models"
	utilities "main/backend/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEntry(username, password, website, notes string) {
	// Encrypt the password we got from the file
	encryptedPassword := utilities.Encrypt([]byte(GetFromMemory().Password), utilities.CreateHash(password))

	// Create a new entry in the database file
	AddToFile(models.Entry{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: string(encryptedPassword),
		Website:  website,
		Notes:    notes,
	}, GetFromMemory().Filename)
}

// return model.Entry or error
func GetEntries() []models.Entry {
	coll := utilities.Client.Database("SeismicSecurity").Collection("entries")
	cursor, err := coll.Find(context.Background(), models.Entry{})
	if err != nil {
		log.Println(
			"Error retrieving entries from database: ",
			err.Error(),
		)
	}

	var entries []models.Entry
	if err = cursor.All(context.Background(), &entries); err != nil {
		log.Println(
			"Error decoding entries from database: ",
			err.Error(),
		)
	}

	return entries
}

func GetEntryByName(username string) models.Entry {
	coll := utilities.Client.Database("SeismicSecurity").Collection("entries")
	var entry models.Entry
	err := coll.FindOne(context.Background(), models.Entry{Username: username}).Decode(&entry)
	if err != nil {
		log.Println(
			"Error retrieving entry from database: ",
			err.Error(),
		)
	}

	return entry
}

func DeleteEntry(id primitive.ObjectID) {
	coll := utilities.Client.Database("SeismicSecurity").Collection("entries")
	_, err := coll.DeleteOne(context.Background(), models.Entry{ID: id})
	if err != nil {
		log.Println(
			"Error deleting entry from database: ",
			err.Error(),
		)
	}

	log.Println("Entry deleted successfully")
}

func UpdateEntry(id primitive.ObjectID, username, password, website, notes string) {
	coll := utilities.Client.Database("SeismicSecurity").Collection("entries")
	_, err := coll.UpdateOne(
		context.Background(),
		models.Entry{ID: id},
		models.Entry{
			Username: username,
			Password: password,
			Website:  website,
			Notes:    notes,
		},
	)
	if err != nil {
		log.Println(
			"Error updating entry in database: ",
			err.Error(),
		)
	}

	log.Println("Entry updated successfully")
}
