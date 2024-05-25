package helpers

import (
	"encoding/json"
	"log"
	"main/backend/models"
	utilities "main/backend/utils"
	"os"
)

func CreateFile(fileName, password string) *error {
	// Create a new file
	file, err := os.Create(fileName + ".Seismic")
	if err != nil {
		return &err
	}

	// Put the encrypted password in the file
	stringToWrite := utilities.Encrypt([]byte(password), utilities.CreateHash(password))
	// put an identifier  before the encrypted password and a 2 new lines after
	stringToWrite = append([]byte("Seismic:"), stringToWrite...)
	stringToWrite = append(stringToWrite, []byte("\n\n")...)
	_, err = file.Write(stringToWrite)
	if err != nil {
		return &err
	}

	StoreInMemory(password, fileName+".Seismic")

	// Close the file
	defer file.Close()

	// Return nil if no error
	return nil
}

func OpenFile(fileName, password string) []byte {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening file")
		return nil
	}

	// Close the file
	defer file.Close()

	// Read the file
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	file.Read(buffer)

	// Get the empty line
	emptyLine := 0
	for i := 0; i < len(buffer); i++ {
		if buffer[i] == 10 && buffer[i+1] == 10 {
			emptyLine = i
			break
		}
	}

	// Get the encrypted password
	encryptedPassword := buffer[8:emptyLine]

	// Decrypt the password
	decryptedPassword := utilities.Decrypt(encryptedPassword, utilities.CreateHash(password))

	// Check if the password is correct
	if string(decryptedPassword) != password {
		log.Println("Password is incorrect")
		return nil
	}

	// Now we filter out the password and the empty line to get the entries
	entries := buffer[emptyLine+2:]

	// decrypt the entries
	entries = utilities.Decrypt(entries, utilities.CreateHash(password))

	StoreInMemory(password, fileName)

	// Return the entries
	return entries
}

func AddToFile(data models.Entry, fileName string) error {
	// Open the file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Close the file
	defer file.Close()

	// Marshal the data
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Encrypt the data
	encryptedData := utilities.Encrypt(dataBytes, utilities.CreateHash(GetFromMemory().Password))

	// Write the data to the file
	_, err = file.Write(encryptedData)
	if err != nil {
		return err
	}

	// Return nil if no error
	return nil
}

var file models.File

func StoreInMemory(password, fileName string) {
	file.Password = password
	file.Filename = fileName
}

func GetFromMemory() models.File {
	return file
}

func GetEntriesFromBytes(entries []byte) ([]models.Entry, error) {
	// Get the entries
	entriesString := string(entries)

	// Split the entries
	entriesArray := utilities.Split(entriesString, "}")

	// Create a slice to store the entries
	var entriesSlice []models.Entry

	// Unmarshal each entry
	for i := 0; i < len(entriesArray); i++ {
		if i == len(entriesArray)-1 {
			break
		}
		// Add the } back to the entry
		entriesArray[i] += "}"
		// Unmarshal the entry
		var entry models.Entry
		err := json.Unmarshal([]byte(entriesArray[i]), &entry)
		if err != nil {
			return nil, err
		}
		// Append the entry to the slice
		entriesSlice = append(entriesSlice, entry)
	}

	// Return the entries
	return entriesSlice, nil
}
