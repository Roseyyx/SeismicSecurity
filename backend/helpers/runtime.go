package helpers

import (
	"log"
	"main/backend/models"
	utilities "main/backend/utils"
)

var UsingFile bool = false

func Setup() {
	if !CheckIfDatabaseExists() {
		// TODO: Add a prompt to create a master password
		UsingFile = true
		CreateMasterPassword("password")
	}
}

func CheckIfDatabaseExists() bool {
	return utilities.HasDatabaseConnection
}

func CreateMasterPassword(password string) {
	File := new(models.File)

	hashedPassword, err := utilities.HashPassword(password)
	if err != nil {
		log.Println(err)
	}

	File.Password = hashedPassword
}
