package helpers

import (
	"log"
	utilities "main/backend/utils"
	"os"
)

var Debug bool = false

func Setup() {
	log.Println("Called Setup")

	// Outdated code
	/*if CheckDocker() && CheckIfDatabaseExists() {
		log.Println("Using Docker")
	}
	if !CheckIfDatabaseExists() {
		// TODO: Add a prompt to create a master password
		UsingFile = true
		CreateMasterPassword("password")
	}*/
}

var FilePath string

func CheckFile() bool {
	if _, err := os.Stat("database.Seismic"); err == nil {
		FilePath = "database.Seismic"
		return true
	}
	return false
}

func CheckIfDatabaseExists() bool {
	return utilities.HasDatabaseConnection
}
