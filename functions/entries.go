package functions

import (
	"context"
	"log"
	"main/models"
	utilities "main/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEntry(username, password, website, notes string) {
	coll := utilities.Client.Database("SeismicSecurity").Collection("entries")
	entry := models.Entry{
		Username: username,
		Password: password,
		Website:  website,
		Notes:    notes,
	}

	_, err := coll.InsertOne(context.TODO(), entry)
	if err != nil {
		log.Println(
			"Error inserting entry into database: ",
			err.Error(),
		)
	}

	log.Println("Entry created successfully")
}

// return model.Entry or error
func GetEntries() []models.Entry {
	coll := utilities.Client.Database("SeismicSecurity").Collection("entries")
	cursor, err := coll.Find(context.Background(), nil)
	if err != nil {
		log.Println(
			"Error retrieving entries from database: ",
			err.Error(),
		)
		return nil
	}

	var entries []models.Entry
	for cursor.Next(context.Background()) {
		var entry models.Entry
		cursor.Decode(&entry)
		entries = append(entries, entry)
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

func CreateAccessEntry(fileName, password string) {
	coll := utilities.Client.Database("SeismicSecurity").Collection("accessEntries")
	entry := models.AccessEntry{
		FileName: fileName,
		Password: password,
	}

	_, err := coll.InsertOne(context.TODO(), entry)
	if err != nil {
		log.Println(
			"Error inserting access entry into database: ",
			err.Error(),
		)
	}

	log.Println("Access entry created successfully")
}
