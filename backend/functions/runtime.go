package functions

import (
	"context"
	"log"
	"main/backend/models"
	utilities "main/backend/utils"
)

func Setup() {
	if !CheckIfDatabaseExists() {
		// TODO: Add a prompt to create a master password
		CreateMasterPassword("password")
	}
}

func CheckIfDatabaseExists() bool {
	// Check the connection
	err := utilities.Client.Ping(context.Background(), nil)

	return err == nil
}

func CreateMasterPassword(password string) {
	File := new(models.File)

	hashedPassword, err := utilities.HashPassword(password)
	if err != nil {
		log.Println(err)
	}

	File.Password = hashedPassword
}
