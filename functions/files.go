package functions

import (
	utilities "main/utils"
	"os"
)

func CreateFile(fileName, password string) *error {
	// Create a new file
	file, err := os.Create(fileName + ".Seismic")
	if err != nil {
		return &err
	}

	// Close the file
	defer file.Close()

	// Encrypt the file
	utilities.EncryptFile(password, file)

	// Return nil if no error
	return nil
}

func OpenFile(fileName, password string) *error {
	// Open the file
	file, err := os.Open(fileName + ".Seismic")
	if err != nil {
		return &err
	}

	// Close the file
	defer file.Close()

	// Decrypt the file
	utilities.DecryptFile(password, file)

	// Return nil if no error
	return nil
}
