package helpers

import (
	"log"
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
	// put an identifier  before the encrypted password
	stringToWrite = append([]byte("Seismic:"), stringToWrite...)
	_, err = file.Write(stringToWrite)
	if err != nil {
		return &err
	}

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

	// Check trough the whole file for the string "Seismic"
	byteIndex := -1
	for i := 0; i < len(buffer); i++ {
		if buffer[i] == 'S' {
			if string(buffer[i:i+8]) == "Seismic:" {
				byteIndex = i + 8
				break
			}
		}
	}

	// If the string "Seismic" is not found, return nil
	if byteIndex == -1 {
		log.Println("File is not a Seismic file")
		return nil
	}

	log.Println(string(buffer[byteIndex:]))

	// Decrypt the password first
	decryptedPassword := utilities.Decrypt(buffer[byteIndex:], utilities.CreateHash(password))
	log.Println(string(decryptedPassword))

	// Compare the decrypted password with the input password
	if string(decryptedPassword) != password {
		log.Println("Password is incorrect")
		return nil
	}

	// Decrypt the rest of the file except the password
	//decryptedData := utilities.Decrypt(buffer[:byteIndex], utilities.CreateHash(password))

	// Return the decrypted data
	return decryptedPassword
}
