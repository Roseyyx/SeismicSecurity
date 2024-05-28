package helpers

import (
	"log"
	"main/backend/models"
	utilities "main/backend/utils"
	"os"
)

func CreateFile(password string) {
	file, err := os.Create("database.Seismic")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	key := utilities.GenerateAESKey()
	hashedPassword := utilities.CreateHash(password)

	// Encrypt the key with the hashed password
	encryptedKey := utilities.Encrypt(key, []byte(hashedPassword))
	log.Println(string(key))

	StoreInMemory(hashedPassword, "database.Seismic", string(encryptedKey))

	// Write the key to the file between the words "STARTKEY" and "ENDKEY"
	_, err = file.WriteString("STARTKEY\n" + string(encryptedKey) + "\nENDKEY")
	if err != nil {
		log.Println(err)
	}

}

func ReadFile(password string) {
	// The key is between the words "STARTKEY" and "ENDKEY"
	file, err := os.Open("database.Seismic")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// Loop through the file and find the string "STARTKEY"
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		log.Println(err)
	}

	// Find the key
	startKey := "STARTKEY\n"
	endKey := "\nENDKEY"

	startIndex := utilities.FindIndex(buffer, []byte(startKey))
	endIndex := utilities.FindIndex(buffer, []byte(endKey))

	// Get the key
	key := buffer[startIndex+len(startKey) : endIndex]

	// Decrypt the key with the password
	hashedPassword := utilities.CreateHash(password)
	decryptedKey := utilities.Decrypt(key, []byte(hashedPassword))

	if decryptedKey == nil {
		log.Println("Incorrect password")
	}

	StoreInMemory(string(hashedPassword), "database.Seismic", string(decryptedKey))
}

var File models.File

func StoreInMemory(password, filename, key string) {
	File.Password = password
	File.Filename = filename
	File.Key = key
}
