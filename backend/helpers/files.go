package helpers

import (
	"encoding/json"
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

	StoreInMemory(hashedPassword, "database.Seismic", string(key))

	// Write the key to the file between the words "STARTKEY" and "ENDKEY"
	_, err = file.WriteString("STARTKEY\n" + string(encryptedKey) + "\nENDKEY")
	if err != nil {
		log.Println(err)
	}

}

func ReadFile(password string) []byte {
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

	// Get the rest of the data in the file after the "ENDKEY"
	data := buffer[endIndex+len(endKey):]

	if len(data) == 0 {
		log.Println("No data in the file")
	}

	// Decrypt the data with the key
	decryptedData := utilities.Decrypt(data, decryptedKey)

	log.Println(string(decryptedData))

	return decryptedData
}

func AddToFile(data models.Entry) {
	file, err := os.OpenFile("database.Seismic", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// transform the data to a byte array and then to a json object
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	// Encrypt the data with the key
	encryptedData := utilities.Encrypt(dataBytes, []byte(File.Key))

	// Write the encrypted data to the file
	_, err = file.Write(encryptedData)
	if err != nil {
		log.Println(err)
	}

}

var File models.File

func StoreInMemory(password, filename, key string) {
	File.Password = password
	File.Filename = filename
	File.Key = key
}

func GetFromMemory() models.File {
	return File
}
