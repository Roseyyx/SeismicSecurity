package helpers

import (
	"log"
	"main/backend/models"
)

func Login(password, file string) {
	receivedBytes := OpenFile(file, password)

	// print the decrypted password
	log.Println(string(receivedBytes))
}

func Register(password string) {
	file := new(models.File)

	file.Filename = "database"

	CreateFile(file.Filename, password)
}
